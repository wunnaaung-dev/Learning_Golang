package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/wunnaaung-dev/payroll-bre/models"
	"github.com/wunnaaung-dev/payroll-bre/services"
	"github.com/wunnaaung-dev/payroll-bre/utils"
)

func GetAllStaffs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	staffs, err := services.GetAllStaffs()

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, staffs, "Staffs retrieved successfully")
}

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var staff models.CreateStaffDTO

	err := json.NewDecoder(r.Body).Decode(&staff)

	if err != nil {
		utils.RespondWithError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdStaff, err := services.InsertStaff(staff)

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithSuccess(w, createdStaff, "Teacher Created Successfully")

}
