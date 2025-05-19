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

func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	teachers, err := services.GetAllTeachers()

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return 
	}

	utils.RespondWithSuccess(w, teachers, "Teachers fetched successfully")

}

func GetTeacherInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	teacher, err := services.GetTeacherInfo(id)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondWithSuccess(w, teacher, "Teacher info fetched successfully")
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var teacher models.CreateTeacherDTO

	err := json.NewDecoder(r.Body).Decode(&teacher)

	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validationErrors := utils.CheckTeacherData(teacher)

	if len(validationErrors) > 0 {
		utils.RespondWithError(w, fmt.Sprintf("Validation failed %s", strings.Join(validationErrors, ", ")), http.StatusBadRequest)
		return
	}

	createdTeacher, err := services.InsertTeacher(teacher)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, createdTeacher, "Teacher Created Successfully")
}

func UpdateTeacherInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.RespondWithError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var teacher models.UpdateTeacherDTO

	err = json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedTeacher, err := services.UpdateTeacher(id, teacher)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, updatedTeacher, "Employee updated successfully")

}
