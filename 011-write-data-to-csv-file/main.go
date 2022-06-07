package main

import(
	"fmt"
	"log"
	"os"
	"encoding/csv"
)

func printFile(s [][]string){
	file,err := os.Create("user.cvs")
	defer file.Close()
	if err != nil{
		log.Fatalln("failed to open file..", err)
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(s)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("success!")
}

func main() {
	data := [][]string{
		{"user_name","ID"},
		{"fua_alfu","1"},
		{"mua_alfu","2"},
		{"smogy_alfu","3"},
	}

	printFile(data)
}