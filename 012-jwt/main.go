package main

import(
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijavala/jwt-go"
)

var(
	key = []byte("key")
)

func createToken()(string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["key"] = time.Now().Add(time.Hour).Unix()
	token, err := token.SignedString(key)
	if err != nil{
		fmt.Println(err.Error())
		return "", err
	}
	return token, nil
}

func testAPI(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "test..")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "home..")
}

func handeler(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api",testAPI)

	fmt.Printf("servering at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}
}

func main(){
	handeler()
}