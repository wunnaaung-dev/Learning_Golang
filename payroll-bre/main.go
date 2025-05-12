package main

import (
	"fmt"
	"github.com/wunnaaung-dev/payroll-bre/router"
	"github.com/wunnaaung-dev/payroll-bre/database"
	"log"
	"net/http"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database connection established successfully")
	r := router.Router()

	fmt.Println("Server starting on the port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
