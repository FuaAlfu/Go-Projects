package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

)

type Post struct {
	Title string
	Body  string
}

func main() {
	
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5454/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	result, err := db.Query("SELECT * FROM posts;")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	// iterate over all the posts and print them
	for result.Next() {
		var p Post
		err := result.Scan(&p.Title, &p.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", p)
	}
}
