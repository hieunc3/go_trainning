package routes

import (
	"go_tranning/book_management/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//Here we register 5 routes mapping URL paths to handlers
	/*
		This is equivalent to how http.HandleFunc() works:
		if an incoming request URL matches one of the paths, the corresponding handler
		is called passing (http.ResponseWriter, *http.Request) as parameters.
	*/
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
