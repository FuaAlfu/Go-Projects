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
func CreateBook(){}

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks,_ := models.GetAllBooks()
	res, _    := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	BookId := vars["book"]
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

func UpdateBook(){}

func DeleteBook(){}