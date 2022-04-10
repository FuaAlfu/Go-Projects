package config

import(
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","fua:fua/simpletest?charset=utf8&pareTime=True&loc=local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}