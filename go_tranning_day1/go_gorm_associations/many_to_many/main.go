package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Uname     string
	Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
	ID   int `gorm:"primary_key"`
	Name string
}
type UserLanguages struct {
	UserId     int
	LanguageId int
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
	Db.Migrator().DropTable(&Language{}, &User{}, &UserLanguages{})
	//Migrate the shema
	Db.AutoMigrate(&Language{}, &User{}, &UserLanguages{})

	langs := []Language{
		{Name: "English"},
		{Name: "Japan"},
	}
	user1 := User{Uname: "John", Languages: langs}
	user2 := User{Uname: "Martin", Languages: langs}
	user3 := User{Uname: "Ray", Languages: langs}

	Db.Create(&user1)
	Db.Create(&user2)
	Db.Create(&user3)

	//Select
	user := &User{}
	Db.Debug().Where("uname=?", "Ray").Find(&user)
	Db.Debug().Model(&user).Association("Languages").Find(&user.Languages)
	fmt.Println(user)
	// Db.Where("uname=?", "Martin").Preload("Languages").Find(&user)
	// fmt.Println(users)

	//Delete
	Db.Debug().Where("uname=?", "Ray").Delete(&user)
	//Db.Select("Languages").Delete(&user)

	//Update
	//Db.Debug().Model(&User{}).Where("uname=?", "Ray").Update("uname", "keran")

}
