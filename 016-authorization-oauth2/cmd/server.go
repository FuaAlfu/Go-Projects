package cmd

import(
	"fmt"
	"net/http"

	"github.com/FuaAlfu/Go-Projects/016-authorization-oauth2/pkg/router"
	"github.com/gorilla/mux"
)

func Connect(){
	fmt.Println("servering..")
	r := mux.NewRouter()
	router.SetupRoutes(r)
	err := http.ListenAndServe(":9090",r)
	if err != nil{
		panic(err)
	}
	
}