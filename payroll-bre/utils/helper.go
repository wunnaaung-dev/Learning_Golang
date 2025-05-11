package utils

import (
	"encoding/json"
	"net/http"

	"github.com/wunnaaung-dev/payroll-bre/models"
)

func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	resp := models.Response[any]{
		Success: false,
		Message: message,
		Data:    nil,
	}
	
	json.NewEncoder(w).Encode(resp)
}

func RespondWithSuccess[T any](w http.ResponseWriter, data T, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	resp := models.Response[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(resp)
}