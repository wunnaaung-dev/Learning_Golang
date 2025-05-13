package services

import (
	"fmt"

	"github.com/wunnaaung-dev/payroll-bre/database"
	"github.com/wunnaaung-dev/payroll-bre/models"
)

func GetAllStaffs() ([]models.StaffResponseDTO, error) {
	db := database.GetDB()

	sqlStatement := `
		SELECT 
			"Employees".id AS staff_id, 
			"Employees".name, 
			"Employees".phone, 
			"Employees".type, 
			"Staffs".role, 
			"Staffs"."maxLeave"
		FROM "Employees"
		RIGHT JOIN "Staffs"
		ON "Employees".id = "Staffs".staff_id
		WHERE "Employees".type = 'Staff';
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var staffs []models.StaffResponseDTO

	for rows.Next() {
		var staff models.StaffResponseDTO
		err := rows.Scan(
			&staff.Staff_ID,
			&staff.Name,
			&staff.Phone,
			&staff.Type,
			&staff.Role,
			&staff.MaxLeave,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		staffs = append(staffs, staff)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return staffs, nil
}

func InsertStaff(staff models.CreateStaffDTO) (models.Staff, error) {
	db := database.GetDB()

	sqlStatement := `
		INSERT INTO "Staffs" (staff_id, role, "maxLeave")
		VALUES
		($1, $2, $3) 
		RETURNING id, staff_id, role, "maxLeave", created_at
	`

	var createdStaff models.Staff

	err := db.QueryRow(sqlStatement, staff.Staff_ID, staff.Role, staff.MaxLeave).Scan(
		&createdStaff.ID,
		&createdStaff.Staff_ID,
		&createdStaff.Role,
		&createdStaff.MaxLeave,
		&createdStaff.CreatedAt,
	)

	if err != nil {
		return models.Staff{}, fmt.Errorf("failed to exectue query: %w", err)
	}

	return createdStaff, nil
}
