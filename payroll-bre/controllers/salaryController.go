package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wunnaaung-dev/payroll-bre/models"
	"github.com/wunnaaung-dev/payroll-bre/services"
	"github.com/wunnaaung-dev/payroll-bre/utils"
)

func GetEmpSalaries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	empType := r.URL.Query().Get("empType")
	if empType == "" {
		utils.RespondWithError(w, "empType query parameter is required", http.StatusBadRequest)
		return
	}

	salaries, err := services.GetSalary(empType)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, salaries, "Salaries fetched successfully")
}

func CreateSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var salary models.CreateSalaryDTO

	err := json.NewDecoder(r.Body).Decode(&salary)

	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdSalary, err := services.InsertSalary(salary)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, createdSalary, "Teacher Created Successfully")

}

func UpdateSalary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var salary models.UpdateSalaryDTO

	err = json.NewDecoder(r.Body).Decode(&salary)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedSalary, err := services.UpdateSalary(id, salary)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, updatedSalary, "Salary updated successfully")
}
