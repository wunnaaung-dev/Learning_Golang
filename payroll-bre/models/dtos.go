package models

import "encoding/json"

type CreateEmployeeDTO struct {
	EmployeeBase
}

type UpdateEmployeeDTO struct {
	ID    int64  `json:"id"`
	Phone string `json:"phone"`
}

type CreateTeacherDTO struct {
	TeacherBase
}

type UpdateTeacherDTO struct {
	Role                    string `json:"role"`
	Total_Classes_Per_Month int    `json:"total_classes_per_month"`
}

type TeacherResponseDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	TeacherBase
}

type CreateStaffDTO struct {
	StaffBase
}

type StaffResponseDTO struct {
	EmployeeBase
	StaffBase
}

type CreateSalaryDTO struct {
	SalaryBase
}

type SalaryResponseDTO struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Monthly_Rate   float64 `json:"monthly_rate"`
	Rate_Per_Class float64 `json:"rate_per_class"`
}

type UpdateSalaryDTO struct {
	Monthly_Rate   float64 `json:"monthly_rate"`
	Rate_Per_Class float64 `json:"rate_per_class"`
}

type UpdateAdjustmentDTO struct {
	Records json.RawMessage `json:"records"`
}

type Result struct {
	Message string
}
