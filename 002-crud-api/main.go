package main

import(
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

type(
	Movie struct{
		ID string `json: "id"`
		Isbn string `json: "isbn`
		Title string `json: "title"`
		Director *Director `json: "director`
	}

	Director struct{
		FirstName string `json: "firstname"`
		LastName string `json: "lastname"`
	}
)

var movies Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request){}
func createMovie(w http.ResponseWriter, r *http.Request){}
func updateMovie(w http.ResponseWriter, r *http.Request){}
func deleteMovie(w http.ResponseWriter, r *http.Request){}

func handler(){
	r := mux.NewRouter()

	movies = append(movies,Movie[ID: "1", Isbn: "4:8227",Title: "Movie Goovie".Director: &Director[FirstName:"john", LastName: "doe"]])
	movies = append(movies,Movie[ID: "2", Isbn: "4:5445",Title: "Movie Dub".Director: &Director[FirstName:"wi", LastName: "wa"]])

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("servering at port 9090\n")
	log.Fatal(http.ListenAndServe(":9090",r))
}

func main() {
	handler()
}