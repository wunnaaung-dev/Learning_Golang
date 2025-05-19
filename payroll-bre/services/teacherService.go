package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"github.com/wunnaaung-dev/payroll-bre/database"
	"github.com/wunnaaung-dev/payroll-bre/models"
)

func GetAllTeachers() ([]models.TeacherResponseDTO, error) {
	db := database.GetDB()

	sqlStatement := `SELECT "Employees".id AS teacher_id, "Employees".name, "Employees".phone, "Teachers".subject, "Teachers".role, "Teachers".total_classes_per_month
						FROM "Employees"
						RIGHT JOIN "Teachers"
						ON "Employees".id = "Teachers".teacher_id;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, fmt.Errorf("unable to execute the query %v", err)
	}

	defer rows.Close()

	var teachers []models.TeacherResponseDTO
	for rows.Next() {
		var teacher models.TeacherResponseDTO
		if err := rows.Scan(&teacher.Teacher_ID, &teacher.Name, &teacher.Phone, &teacher.Subject, &teacher.Role, &teacher.Total_Classes_Per_Month); err != nil {
			return nil, fmt.Errorf("unable to scan the row: %v", err)
		}
		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return teachers, nil
}

func GetTeacherInfo(id int) (models.TeacherResponseDTO, error) {
	db := database.GetDB()

	sqlStatement := `
		SELECT "Employees".id AS teacher_id, "Employees".name, "Employees".phone, "Teachers".subject, "Teachers".role, "Teachers".total_classes_per_month
		FROM "Employees"
		RIGHT JOIN "Teachers"
		ON "Employees".id = "Teachers".teacher_id
		WHERE "Teachers".teacher_id = $1;
	`

	var teacher models.TeacherResponseDTO
	err := db.QueryRow(sqlStatement, id).Scan(
		&teacher.Teacher_ID,
		&teacher.Name,
		&teacher.Phone,
		&teacher.Subject,
		&teacher.Role,
		&teacher.Total_Classes_Per_Month,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TeacherResponseDTO{}, fmt.Errorf("teacher with ID %d not found", id)
		}
		return models.TeacherResponseDTO{}, fmt.Errorf("unable to execute the query: %v", err)
	}

	return teacher, nil
}

func InsertTeacher(teacher models.CreateTeacherDTO) (models.Teacher, error) {
	db := database.GetDB()

	sqlStatement := `
		INSERT INTO "Teachers" (teacher_id, subject, role, total_classes_per_month)
		VALUES
		($1, $2, $3, $4) 
		RETURNING id, created_at, teacher_id, subject, role, total_classes_per_month
	`
	var createdTeacher models.Teacher
	err := db.QueryRow(sqlStatement, teacher.Teacher_ID, teacher.Subject, teacher.Role, teacher.Total_Classes_Per_Month).Scan(
		&createdTeacher.ID,
		&createdTeacher.CreatedAt,
		&createdTeacher.Teacher_ID,
		&createdTeacher.Subject,
		&createdTeacher.Role,
		&createdTeacher.Total_Classes_Per_Month,
	)

	if err != nil {
		return models.Teacher{}, fmt.Errorf("failed to exectue query: %w", err)
	}

	return createdTeacher, nil
}

func UpdateTeacher(id int, teacher models.UpdateTeacherDTO) (models.Teacher, error) {
	db := database.GetDB()

	// Verify the teacher exists
	var exists bool
	if err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "Teachers" WHERE teacher_id = $1)`, id).Scan(&exists); err != nil {
		return models.Teacher{}, fmt.Errorf("database error while checking teacher existence: %w", err)
	}
	if !exists {
		return models.Teacher{}, fmt.Errorf("teacher with ID %d not found", id)
	}

	// Build dynamic update query
	setClauses, args := []string{}, []any{}
	if teacher.Role != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"role" = $%d`, len(args)+1))
		args = append(args, teacher.Role)
	}
	if teacher.Total_Classes_Per_Month != 0 {
		setClauses = append(setClauses, fmt.Sprintf(`"total_classes_per_month" = $%d`, len(args)+1))
		args = append(args, teacher.Total_Classes_Per_Month)
	}
	
	if len(setClauses) == 0 {
		return models.Teacher{}, fmt.Errorf("no fields to update")
	}

	// Execute update query
	sqlUpdate := fmt.Sprintf(`
		UPDATE "Teachers"
		SET %s
		WHERE teacher_id = $%d
		RETURNING id, created_at, teacher_id, subject, role, total_classes_per_month;
	`, strings.Join(setClauses, ", "), len(args)+1)
	args = append(args, id)

	var updatedTeacher models.Teacher
	
	if err := db.QueryRow(sqlUpdate, args...).Scan(
		&updatedTeacher.ID,
		&updatedTeacher.CreatedAt,
		&updatedTeacher.Teacher_ID,
		&updatedTeacher.Subject,
		&updatedTeacher.Role,
		&updatedTeacher.Total_Classes_Per_Month,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Teacher{}, fmt.Errorf("teacher with ID %d exists but update failed", id)
		}
		return models.Teacher{}, fmt.Errorf("failed to execute update query: %w", err)
	}

	return updatedTeacher, nil
}
