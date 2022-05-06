package main

import (
	"fmt"
	"go_tranning/go_crud/utils"
	"log"
)

type Book struct {
	Id               int
	Name             string
	Price            float32
	Publication_year int
	Author           string
}

func main() {
	//getAllBooks()

	// book1 := Book{Name: "Don Quixote", Price: 144.5, Publication_year: 1998, Author: "Miguel de Cervantes"}
	// insertBookToDb(book1)

	//getBookById(4)

	//updateBook(6, Book{Name: "Beloved", Price: 184.5, Publication_year: 1873, Author: "Toni Morrison"})

	//deleteBook(6)
}

func getAllBooks() {
	res, err := utils.GetConnection().Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	fmt.Printf("%s %15s %25s %14s %5s\n", "ID", "Name", "Price", "Year", "Author")

	for res.Next() {
		var book Book
		err := res.Scan(&book.Id, &book.Name, &book.Price, &book.Publication_year, &book.Author)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d %30s %15f %10d %5s\n", book.Id, book.Name, book.Price, book.Publication_year, book.Author)
	}
}

func getBookById(id int) {
	sqlQuery := "SELECT id, name,price,publication_year,author from books where id = ?"
	stmt, err := utils.GetConnection().Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	var book Book
	res, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	if res.Next() {
		res.Scan(&book.Id, &book.Name, &book.Price, &book.Publication_year, &book.Author)
	}
	fmt.Println(book)
}

func insertBookToDb(book Book) {
	sqlQuery := "INSERT INTO books(name,price,publication_year,author) VALUES(?,?,?,?)"
	stmt, err1 := utils.GetConnection().Prepare(sqlQuery)

	if err1 != nil {
		log.Fatal(err1)
	}
	res, err2 := stmt.Exec(book.Name, book.Price, book.Publication_year, book.Author)
	if err2 != nil {
		log.Fatal(err2)
	}
	rowsAffected, err3 := res.RowsAffected()
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(rowsAffected)
}

func updateBook(id int, book Book) {
	sqlQuery := "UPDATE books SET name=?,price=?,publication_year=?,author=? where id= ?"
	stmt, err := utils.GetConnection().Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(book.Name, book.Price, book.Publication_year, book.Author, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func deleteBook(id int) {
	sqlQuery := "DELETE FROM books WHERE id = ?"
	stmt, err := utils.GetConnection().Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rowsAffected)
}
