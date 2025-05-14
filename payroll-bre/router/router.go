package router

import (
	"github.com/wunnaaung-dev/payroll-bre/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// API Endpoints for employees
	router.HandleFunc("/api/employees", controllers.GetAllEmployees).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employees", controllers.CreateEmp).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", controllers.UpdateEmployee).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", controllers.DeleteEmployee).Methods("DELETE", "OPTIONS")

	// API Endpoints for teachers
	router.HandleFunc("/api/teachers", controllers.GetAllTeachers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/teachers", controllers.CreateTeacher).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/teacher/{id}", controllers.UpdateTeacherInfo).Methods("PATCH", "OPTIONS")

	// API Endponts for staffs
	router.HandleFunc("/api/staffs", controllers.GetAllStaffs).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/staffs", controllers.CreateStaff).Methods("POST", "OPTIONS")

	// API Endpoints for salary
	router.HandleFunc("/api/salary", controllers.CreateSalary).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/salary", controllers.GetEmpSalaries).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/salary/{id}", controllers.UpdateSalary).Methods("PATCH", "OPTIONS")

	return router
}
