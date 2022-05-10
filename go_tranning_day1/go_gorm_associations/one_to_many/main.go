package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	username = "root"
	password = "12345678"
	hostname = "127.0.0.1:3306"
)

type Customer struct {
	CustomerID   int `gorm:"primary_key"`
	CustomerName string
	Contacts     []Contact `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:CustId"`
}

type Contact struct {
	ContactID   int `gorm:"primary_key"`
	CountryCode int
	MobileNo    uint
	CustId      int
}

func dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)
}

var Db *gorm.DB

func main() {
	Db, _ := gorm.Open(mysql.Open(dsn("associations")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})

	//Drop table if exists
	Db.Migrator().DropTable(&Contact{}, &Customer{})
	//Migrate the shema
	Db.AutoMigrate(&Contact{}, &Customer{})

	Custs1 := Customer{CustomerName: "Hieu", Contacts: []Contact{
		{CountryCode: 91, MobileNo: 956112},
		{CountryCode: 91, MobileNo: 997555},
	}}
	Custs2 := Customer{CustomerName: "Huy", Contacts: []Contact{
		{CountryCode: 90, MobileNo: 808988},
		{CountryCode: 90, MobileNo: 909699},
	}}
	Custs3 := Customer{CustomerName: "Quang", Contacts: []Contact{
		{CountryCode: 75, MobileNo: 798088},
		{CountryCode: 75, MobileNo: 965755},
	}}
	Custs4 := Customer{CustomerName: "Minh", Contacts: []Contact{
		{CountryCode: 81, MobileNo: 805510},
		{CountryCode: 81, MobileNo: 758863},
	}}

	//Insert
	Db.Create(&Custs1)
	Db.Create(&Custs2)
	Db.Create(&Custs3)
	Db.Create(&Custs4)

	//Select
	customers := &Customer{}
	customers1 := &Customer{}
	Db.Where("customer_name=?", "Huy").Preload("Contacts").Find(&customers) //get data & their associations

	Db.Where("customer_name=?", "Minh").Preload("Contacts").Find(&customers1) //get data & their associations
	fmt.Println(customers)
	fmt.Println(customers1)

	//Update
	//contacts := &Contact{}
	Db.Model(&Contact{}).Where("cust_id=?", 2).Update("country_code", 77)

	//Delete
	Db.Where("customer_name=?", customers.CustomerName).Delete(&customers)
	//Db.Select("Contacts").Where("customer_name=?", customers.CustomerName).Delete(&customers)

}
