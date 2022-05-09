package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

const (
	username = "root"
	password = "12345678"
	hostname = "127.0.0.1:3306"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", username, password, hostname, dbname)
}

func Connect() {
	d, err := gorm.Open(mysql.Open(dsn("store_book")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
