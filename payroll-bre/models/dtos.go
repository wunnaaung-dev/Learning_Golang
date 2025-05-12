package models

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
