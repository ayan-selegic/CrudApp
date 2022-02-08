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
const db_url = "b64dd291402401:35ee02b4@tcp(us-cdbr-east-05.cleardb.net)/heroku_1ee2768dc3824da"

// database connection
func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB") // exception that arises as runtime
	}
	DB.AutoMigrate(&User{})
}
