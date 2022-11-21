package database

import(
	"log"
	"os"

	"gorm.io/gorm"
)

type DBInstance struct {
	db *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil{
		log.Fatal("fatal: Failed to connect to the database! \n",err.Error())
		os.Exit(2)
	}
	log.Printf("Connected to the database")
	db.Logger = Logger.Defualt.LogMode(Logger.Info)
	log.Println("running migrations.. ")
}