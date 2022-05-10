package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Place struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Town   Town
	TownId int `gorm:"ForeignKey:id"`
}

type Town struct {
	ID   int `gorm:"primary_key"`
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

var Db *gorm.DB

func main() {
	Db, _ := gorm.Open(mysql.Open(dsn("associations")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})

	//Drop table if exists
	Db.Migrator().DropTable(&Place{}, &Town{})
	//Migrate the shema
	Db.AutoMigrate(&Place{}, &Town{})

	t1 := Town{
		Name: "Pune",
	}
	t2 := Town{
		Name: "Mumbai",
	}
	t3 := Town{
		Name: "Hyderabad",
	}

	p1 := Place{
		Name: "Katraj",
		Town: t1,
	}
	p2 := Place{
		Name: "Thane",
		Town: t2,
	}
	p3 := Place{
		Name: "Secundarabad",
		Town: t3,
	}
	p4 := Place{
		Name: "Secundarabad",
		Town: t1,
	}
	/*

	 */
	Db.Create(&p1)
	Db.Create(&p2)
	Db.Create(&p3)
	Db.Create(&p4)

	// Start Association Mode
	//var place Place
	//var town Town
	// `place` is the source model, it must contains primary key
	// `Town` is a relationship's field name
	// If the above two requirements matched, the AssociationMode should be started successfully, or it should return error
	//err := Db.Model(&t1).Association("Town").Error

	//Delete
	//Db.Select("Place").Where("name=?", "Hyderabad").Delete(&t3) // delete town''s place when deleting town
	//Db.Select("Town").Delete(&p3) // delete place's Town relations when deleting place
	//Db.Where("name=?", "Secundarabad").Delete(&Place{})

	//Update
	Db.Model(&Place{}).Where("id=?", 1).Update("name", "Shivaji Nagar")

	//Select
	places := Place{}
	Db.Where("name=?", "Shivaji Nagar").Find(&places)
	// err := Db.Model(&places).Association("Town").Find(&places.Town)
	// fmt.Println(err)

}
