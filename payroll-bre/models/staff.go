package models

import "time"

type Staff struct {
	ID        int64     `json:"id"`
	StaffBase
	CreatedAt time.Time `json:"-"`
}

type StaffBase struct {
	Staff_ID int64  `json:"staff_id"`
	Role     string `json:"role"`
	MaxLeave int `json:"maxLeave"`
}
