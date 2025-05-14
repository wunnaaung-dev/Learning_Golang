package models

import "time"

type Salary struct {
	ID int64 `json:"id"`
	SalaryBase
	CreatedAt time.Time `json:"-"`
}

type SalaryBase struct {
	Employee_ID    int64   `json:"employee_id"`
	Monthly_Rate   float64 `json:"monthly_rate"`
	Rate_Per_Class float64 `json:"rate_per_class"`
}
