package cmd

import(
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request){
	w.Header([]byte("run\n"))
}

func Handler(){
	http.HandleFunc("/",mainPage)
	fmt.Printf("servering at port 9090\n")
	if err := http.ListenAndServe(":9090",nil); err != nil{
		log.Fatal(err)
	}
}

func Test(){
	fmt.Println("test..")
}