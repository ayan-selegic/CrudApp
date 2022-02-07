package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // converting this struct into object model via ORM
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
}

var DB *gorm.DB
var err error

// "username:password@protocol(ip)/local_dbname"
const db_url = "root:BOMB@y26@tcp(127.0.0.1:3306)/godb"

// database connection
func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB") // exception that arises as runtime
	}
	DB.AutoMigrate(&User{})
}
