package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

// Initialize database connection
func init() {
	// Load env file, panic if unavailable
	if envErr := godotenv.Load(); envErr != nil {
		log.Panic(envErr)
	}

	// Store env in variables
	dbUser, dbPwd, dbName := os.Getenv("db_user"), os.Getenv("db_pwd"), os.Getenv("db_name")
	dbHost, dbPort := os.Getenv("db_host"), os.Getenv("db_port")

	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
	log.Println("Connecting to DB " + dbName + " as user " + dbUser)

	// Connect to DB
	var dbErr error
	db, dbErr = gorm.Open("mysql", dbUri)
	if dbErr != nil {
		log.Panic(dbErr.Error())
	}

	db.AutoMigrate(&User{})
}

// Get database instance
func GetDB() *gorm.DB {
	return db
}
