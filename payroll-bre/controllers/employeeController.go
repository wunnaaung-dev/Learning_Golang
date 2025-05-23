package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/wunnaaung-dev/payroll-bre/models"
	"github.com/wunnaaung-dev/payroll-bre/services"
	"github.com/wunnaaung-dev/payroll-bre/utils"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := services.GetAllEmployees()

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondWithSuccess(w, employees, "Employees fetched successfully")
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	employee, err := services.GetEmployeeInfo(id)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondWithSuccess(w, employee, "Employee fetched successfully")
}

func CreateEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var employee models.CreateEmployeeDTO

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validationErrors := utils.ValidateEmployeeData(employee)
	if len(validationErrors) > 0 {
		utils.RespondWithError(w, fmt.Sprintf("Validation failed %s", strings.Join(validationErrors, ", ")), http.StatusBadRequest)
		return
	}

	createdEmployee, err := services.InsertEmployee(employee)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, createdEmployee, "Employee created successfully")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee models.UpdateEmployeeDTO

	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validationErrors := utils.CheckEmployeePhone(employee)

	if len(validationErrors) > 0 {
		utils.RespondWithError(w, fmt.Sprintf("Validation failed %s", strings.Join(validationErrors, ", ")), http.StatusBadRequest)
		return
	}

	employee.ID = int64(id)
	updatedEmployee, err := services.UpdateEmployee(employee)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, updatedEmployee, "Employee updated successfully")
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteEmployee(id)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, struct{}{}, "Deleted Successfully")
}
