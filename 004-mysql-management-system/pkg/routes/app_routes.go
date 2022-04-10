package routes

import(
	"fmt"
	"github.com/FuaAlfu/Go-Projects/004-mysql-management-system/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}