package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "3000"
	SERVER_TYPE = "tcp"

	username = "root"
	password = "12345678"
	hostname = "127.0.0.1:3306"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", username, password, hostname, dbname)
}

type Info struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

/*
	Socket server in Golang, concurrently handles clients, receives messages from each and return the
	message back in reply
*/
func main() {

	d, err := gorm.Open(mysql.Open(dsn("store_book")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})

	if err != nil {
		panic(err)
	}
	//Drop table if exists
	//db.Migrator().DropTable(&Product{}, &Category{})
	d.Migrator().DropTable(&Info{})

	//Migrate the shema
	d.AutoMigrate(&Info{})

	fmt.Println("Server Running...")
	// Start the server and listen for incoming connections.
	ser, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer ser.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		c, err := ser.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected")

		// Print client remote connection network address
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go hanleConnection(c, d)
	}
}

// handleConnection handles logic for a single connection request.
func hanleConnection(conn net.Conn, db *gorm.DB) {
	//Buffer client input until a newline
	buffer, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Client left ", conn.RemoteAddr().String())
		conn.Close()
		return
	}

	// Print response message, stripping newline character.
	log.Println("Client message: ", string(buffer[:len(buffer)-1]))

	data := Info{}
	json.Unmarshal([]byte(buffer), &data)
	db.Create(&data)

	// Send response message to the client.
	//conn.Write(buffer)

	// Restart the process.
	hanleConnection(conn, db)
}
