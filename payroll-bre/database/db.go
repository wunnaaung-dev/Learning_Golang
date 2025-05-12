package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbConn *sql.DB
	dbConn, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		return err
	}

	dbConn.SetMaxOpenConns(25)
	dbConn.SetMaxIdleConns(5)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	if err = dbConn.Ping(); err != nil {
		return err
	}

	// Set the global db variable
	db = dbConn

	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		log.Panic("Database connection not initialized. Call InitDB() first.")
	}
	
	return db
}
