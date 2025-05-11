package models
import "time"

type Employee struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	EmployeeBase
	IsWorking bool `json:"isWorking"`
}

type CreateEmployeeDTO struct {
	EmployeeBase
}

type UpdateEmployeeDTO struct {
	ID int64 `json:"id"`
	Phone string `json:"phone"`
}

type EmployeeBase struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Type  string `json:"type"`
}