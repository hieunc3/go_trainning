package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id        int64
	Birthday  time.Time
	Age       int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Emails            []Email       // One-To-Many relationship (has many)
	BillingAddress    Address       // One-To-One relationship (has one)
	BillingAddressId  sql.NullInt64 // Foreign key of BillingAddress
	ShippingAddress   Address       // One-To-One relationship (has one)
	ShippingAddressId int64         // Foreign key of ShippingAddress
	IgnoreMe          int64         `sql:"-"`                          // Ignore this field
	Languages         []Language    `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}
type Email struct {
	Id         int64
	UserId     int64  // Foreign key for User (belongs to)
	Email      string `sql:"type:varchar(100);"` // Set field's type
	Subscribed bool
}
type Address struct {
	Id       int64
	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `sql:"type:varchar(100)"`
	Post     sql.NullString `sql:"not null"`
}

type Language struct {
	Id   int64
	Name string
}

const (
	username = "root"
	password = "12345678"
	hostname = "127.0.0.1:3306"
)

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)
}
func main() {
	db, _ := gorm.Open(mysql.Open(dsn("associations")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})

	//Drop table if exists
	db.Migrator().DropTable(&Language{}, &User{}, &Email{}, &Address{})
	//Migrate the shema
	db.AutoMigrate(&Language{}, &User{}, &Email{}, &Address{})

	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails:          []Email{{Email: "jinzhu@example.com"}, {Email: "jinzhu-2@example@example.com"}},
		Languages:       []Language{{Name: "ZH"}, {Name: "EN"}},
	}

	db.Create(&user)

	//select
	var u User
	db.Where("name=?", "jinzhu").Find(&u)
	db.Model(&user).Association("Emails").Find(&u.Emails)
	// fmt.Println(u.Emails)
	fmt.Printf("%+v", u)

	//var user1 User
	var l Language
	codes := []string{"zh-CN", "en-US", "ja-JP"}
	db.Model(&user).Where("code IN ?", codes).Association("Languages").Find(&l)
	fmt.Println(l)
}
