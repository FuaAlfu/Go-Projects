package main

import(
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = defaultPort
	}
}