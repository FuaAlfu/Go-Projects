package cmd

import(
	"fmt"
	"log"
	
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/FuaAlfu/Go-Projects/004-mysql-management-system/pkg/routes"
)

func handler(){
	fmt.Println("Listening at port: 8080")
	
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8080",r))
}
