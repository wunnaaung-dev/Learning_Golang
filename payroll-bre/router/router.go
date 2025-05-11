package router

import (
	"github.com/wunnaaung-dev/payroll-bre/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/employees", controllers.GetAllEmployees).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/employees", controllers.CreateEmp).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/employee/{id}", controllers.UpdateEmployee).Methods("PATCH", "OPTIONS")

	router.HandleFunc("/api/employee/{id}", controllers.DeleteEmployee).Methods("DELETE", "OPTIONS")

	return router
}