package main

import (
	"go_tranning/book_management/pkg/models"
	"go_tranning/book_management/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	var listBooks = []models.Book{
		{Name: " In Search of Lost Time", Author: "Marcel Proust", Publication: "1989"},
		{Name: "Ulysses", Author: "James Joyce", Publication: "1999"},
		{Name: "Don Quixote ", Author: "Miguel de Cervantes", Publication: "1994"},
		{Name: "One Hundred Years of Solitude", Author: "Gabriel Garcia Marquez", Publication: "1892"},
	}
	for _, v := range listBooks {
		v.CreateBook()
	}
}
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
