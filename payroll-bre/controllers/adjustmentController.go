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

func TestingBonusRule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	rule := services.TestingRule(id)

	utils.RespondWithSuccess(w, rule, "Success")
}

func GetEmployeeSalaryAdjustment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	adjustment, err := services.GetSalaryAdjustment(id)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondWithSuccess(w, adjustment, "Employee Salary Adjustment fetched successfully")
}

func CreateEmployeeSalaryAdjustment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var adjustment models.CreateAdjustmentDTO

	err := json.NewDecoder(r.Body).Decode(&adjustment)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdAdjustment, err := services.CreateSalaryAdjustment(adjustment)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, createdAdjustment, "Employee Salary Adjustment created successfully")
}

func UpdateEmployeeSalaryAdjustment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var adjustment models.UpdateAdjustmentDTO
	err = json.NewDecoder(r.Body).Decode(&adjustment)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedAdjustment, err := services.UpdateSalaryAdjustment(id, adjustment)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, updatedAdjustment, "Employee Salary Adjustment updated successfully")
}
