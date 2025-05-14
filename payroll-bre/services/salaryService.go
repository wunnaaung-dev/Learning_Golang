package services

import (
	"fmt"
	"strings"

	"github.com/wunnaaung-dev/payroll-bre/database"
	"github.com/wunnaaung-dev/payroll-bre/models"
)

func InsertSalary(salary models.CreateSalaryDTO) (models.Salary, error) {
	db := database.GetDB()

	sqlStatement := `
		INSERT INTO "Salaries" (employee_id, monthly_rate, rate_per_class)
		VALUES
		($1, $2, $3)
		RETURNING id, employee_id, monthly_rate, rate_per_class, created_at
	`

	var createdSalary models.Salary

	err := db.QueryRow(sqlStatement, salary.Employee_ID, salary.Monthly_Rate, salary.Rate_Per_Class).Scan(
		&createdSalary.ID,
		&createdSalary.Employee_ID,
		&createdSalary.Monthly_Rate,
		&createdSalary.Rate_Per_Class,
		&createdSalary.CreatedAt,
	)

	if err != nil {
		return models.Salary{}, fmt.Errorf("failed to exectue query: %w", err)
	}

	return createdSalary, nil
}

func GetSalary(empType string) ([]models.SalaryResponseDTO, error) {
	db := database.GetDB()

	sqlStatement := `
		SELECT "Employees".id, "Employees".name, "Employees".type, "Salaries".monthly_rate, "Salaries".rate_per_class
		FROM "Employees"
		RIGHT JOIN "Salaries"
		ON "Employees".id = "Salaries".employee_id
		WHERE "Employees".type = $1
	`

	rows, err := db.Query(sqlStatement, empType)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var salaries []models.SalaryResponseDTO

	for rows.Next() {
		var salary models.SalaryResponseDTO
		err := rows.Scan(
			&salary.ID,
			&salary.Name,
			&salary.Type,
			&salary.Monthly_Rate,
			&salary.Rate_Per_Class,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		salaries = append(salaries, salary)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return salaries, nil

}

func UpdateSalary(id int, salaryDTO models.UpdateSalaryDTO) (models.Salary, error) {
	db := database.GetDB()

	var exists bool
	if err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM "Salaries" WHERE employee_id = $1)`, id).Scan(&exists); err != nil {
		return models.Salary{}, fmt.Errorf("database error while checking salary existence: %w", err)
	}
	if !exists {
		return models.Salary{}, fmt.Errorf("salary with ID %d not found", id)
	}

	setClauese, args := []string{}, []any{}
	if salaryDTO.Monthly_Rate != 0 {
		setClauese = append(setClauese, fmt.Sprintf(`"monthly_rate" = $%d`, len(args)+1))
		args = append(args, salaryDTO.Monthly_Rate)
	}
	if salaryDTO.Rate_Per_Class != 0 {
		setClauese = append(setClauese, fmt.Sprintf(`"rate_per_class" = $%d`, len(args)+1))
		args = append(args, salaryDTO.Rate_Per_Class)
	}

	if len(setClauese) == 0 {
		return models.Salary{}, fmt.Errorf("no fields to update")
	}

	sqlStatement := fmt.Sprintf(
		`
		UPDATE "Salaries"
		SET %s
		WHERE employee_id = $%d
		RETURNING id, employee_id, monthly_rate, rate_per_class, created_at
	`, strings.Join(setClauese, ", "), len(args)+1)
	
	args = append(args, id)
	
	var updatedSalary models.Salary

	err := db.QueryRow(sqlStatement, args...).Scan(
		&updatedSalary.ID,
		&updatedSalary.Employee_ID,
		&updatedSalary.Monthly_Rate,
		&updatedSalary.Rate_Per_Class,
		&updatedSalary.CreatedAt,
	)

	if err != nil {
		return models.Salary{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return updatedSalary, nil
}
