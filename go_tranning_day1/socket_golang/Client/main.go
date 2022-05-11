package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "3000"
	SERVER_TYPE = "tcp"
)

type Doc struct {
}

/*
	Socket client in Golang, prompts for and transfers a message to the srver and displays the returned reply
*/
func main() {

	//establish connection
	fmt.Println("Connecting to " + SERVER_TYPE + " server " + SERVER_HOST + ":" + SERVER_PORT)

	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	//reader := bufio.NewReader(os.Stdin)

	// run loop forever, until exit.
	// for {
	//Prompting message.
	//fmt.Println("Text tp send: ")

	// Read in input until newline, Enter key.
	//input, _ := reader.ReadString('\n')

	// Send to file
	f, err := os.Open("../text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		input := scan.Text() + "\n"
		time.Sleep(1 * time.Second)
		connection.Write([]byte(input))
	}

	// Send to socket connection.
	//connection.Write([]byte(input))

	// Listen for relay.
	//message, _ := bufio.NewReader(connection).ReadString('\n')
	// Print server relay.
	//log.Print("Server relay:", message)

	// }
}
