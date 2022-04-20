package cmd

import(
	"fmt"
	"net/http"

	"github.com/FuaAlfu/Go-Projects/006-server/pkg/router"
	"github.com/gorilla/mux"
)

func Servering(){
	fmt.Println("servering..")
	r := mux.NewRouter()
	router.SetupRoutes(r)
	err := http.ListenAndServe(":9090",r)
	if err != nil{
		panic(err)
	}
	
}