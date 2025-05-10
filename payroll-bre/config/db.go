package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Printf("Warning: Error loading .env.local file: %v\n", err)
	}

	connStr := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Printf("Error opening DB: %v\n", err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("❌ Connection failed: %v\n", err)
		return
	}

	fmt.Println("✅ Connected to Supabase PostgreSQL successfully!")

}
