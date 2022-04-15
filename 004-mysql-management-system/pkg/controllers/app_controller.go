package controllers

import(
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	

	"github.com/gorilla/mux"
	"github.com/FuaAlfu/Go-Projects/004-mysql-management-system/pkg/models"
	"github.com/FuaAlfu/Go-Projects/004-mysql-management-system/pkg/utils"
)
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r.CreateBook)
	b := CreateBook.CreateBook()
	res, _    := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	newBooks,_ := models.GetAllBooks()
	res, _    := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing..")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var uodateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing..")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing..")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}