package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	username = "root"
	password = "12345678"
	hostname = "127.0.0.1:3306"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func InitializeMySQL() {
	dbConnection, err := sql.Open("mysql", dsn("bookstore"))
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = dbConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	db = dbConnection
	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetMaxIdleConns(5)
	dbConnection.SetConnMaxLifetime(time.Second * 10)
}

func GetConnection() *sql.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}

func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}
