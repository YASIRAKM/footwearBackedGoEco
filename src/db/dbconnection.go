package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	// Replace with your MySQL credentials and database name
	dsn := "root:IBfrshdvUnKVianbererJIMQfgmFrfew@tcp(junction.proxy.rlwy.net:53264)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open the database:", err)
	}

	// Test the database connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Database connected successfully")
}

// CloseDB closes the database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatal("Failed to close the database:", err)
	}
}
