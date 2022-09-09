package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db *gorm.DB
)
var password string = ""
var user string = "root"

func Connect() {
	d, err := gorm.Open("mysql", user+":"+password+"@tcp(127.0.0.1:3306)/backend_desarrollo_web?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func setVariablesDBDefault() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	password = os.Getenv("PASSWORD_DB")
	user = os.Getenv("USER_DB")
}

func GetDB() *gorm.DB {
	return db
}
