package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*

 */
type Product struct {
	gorm.Model
	Code       string
	Price      uint
	CategoryID uint
	Category   Category
}

type Category struct {
	gorm.Model
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
	db, err := gorm.Open(mysql.Open(dsn("store")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})
	if err != nil {
		panic("failed to connect database")
	}

	//Drop table if exists
	db.Migrator().DropTable(&Product{})
	db.Migrator().DropTable(&Category{})

	//Migrate the shema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})

	var categoryList = []Category{
		{Name: "vegestable"},
		{Name: "meat"},
		{Name: "seafood"},
	}
	db.Create(categoryList)

	//Create record
	product := Product{Code: "A321", Price: 100, CategoryID: 1}
	// var category Category
	db.Create(&product)

	//batch insert
	var productList = []Product{
		{Code: "A373", Price: 102, CategoryID: 2},
		{Code: "A374", Price: 106, CategoryID: 3},
		{Code: "A442", Price: 116, CategoryID: 1},
		{Code: "A307", Price: 156, CategoryID: 3},
	}
	db.Create(productList)

	//Read records
	getAllProducts(db)
	getAllCategories(db)

	//Get product by id
	getProductById(db, 1)

	//Update record
	updateProductById(db, 1)

	//Delete product
	deleteProductById(db, 1)
}

// func formatDateTime(t time.Time) string {
// 	return t.Format("2006-January-02")
// }

func getAllProducts(db *gorm.DB) {
	fmt.Println("List Products")
	var products []Product
	res := db.Find(&products)
	if res.Error != nil {
		fmt.Println("Err")
	} else {
		for i := 0; i < len(products); i++ {
			fmt.Println(products[i])
		}
	}
}
func getProductById(db *gorm.DB, id int) {
	var product Product
	db.First(&product, id) //find product by id
	fmt.Printf("%v\n", product)
	fmt.Printf("%+v\n", product.CreatedAt.Format(time.ANSIC))
}
func updateProductById(db *gorm.DB, id int) {
	resutl := db.Model(&Product{}).Where("id", id).Update("Price", 120) //update product with id = 1
	if resutl.RowsAffected != 1 {
		fmt.Println("Update Error")
	} else {
		fmt.Println("Update Successfully!")
	}
}

func deleteProductById(db *gorm.DB, id int) {
	resutl := db.Delete(&Product{}, 3)
	if resutl.RowsAffected != 1 {
		fmt.Println("Delete Error")
	} else {
		fmt.Println("Delete Successfully")
	}
}

func getAllCategories(db *gorm.DB) {
	fmt.Println("List Categories")
	var categories []Category
	res1 := db.Find(&categories)
	if res1.Error != nil {
		fmt.Println("Err")
	} else {
		for i := 0; i < len(categories); i++ {
			fmt.Println(categories[i])
		}
	}
}
