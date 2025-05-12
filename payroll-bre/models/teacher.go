package models

import "time"

type Teacher struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"-"`
	TeacherBase
}

type TeacherBase struct {
	Teacher_ID int64 `json:"teacher_id"`
	Subject string `json:"subject"`
	Role string `json:"role"`
	Total_Classes_Per_Month int `json:"total_classes_per_month"`
}