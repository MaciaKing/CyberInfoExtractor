package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbport := os.Getenv("POSTGRES_PORT")
	// sslmode := os.Getenv("SSLMODE")
	sslmode := "disable"
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + dbport + " sslmode=" + sslmode

	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection successful.")
}

func Migrate() {

}
