package services

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wunnaaung-dev/payroll-bre/database"
	"github.com/wunnaaung-dev/payroll-bre/models"
)

func GetAllEmployees() ([]models.Employee, error) {
	db := database.GetDB()

	sqlStatement := `SELECT id, created_at, name, type, phone, "isWorking" FROM "Employees" WHERE "isWorking" = true`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the query: %v", err)
	}

	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.CreatedAt, &employee.Name, &employee.Type, &employee.Phone, &employee.IsWorking); err != nil {
			return nil, fmt.Errorf("unable to scan the row: %v", err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return employees, nil
}

func InsertEmployee(employee models.CreateEmployeeDTO) (models.Employee, error) {
	db := database.GetDB()

	sqlStatement := `
		INSERT INTO "Employees" (name, phone, type)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, name, type, phone, "isWorking"
	`

	var createdEmployee models.Employee
	err := db.QueryRow(sqlStatement, employee.Name, employee.Phone, employee.Type).Scan(
		&createdEmployee.ID,
		&createdEmployee.CreatedAt,
		&createdEmployee.Name,
		&createdEmployee.Type,
		&createdEmployee.Phone,
		&createdEmployee.IsWorking,
	)
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to insert employee: %w", err)
	}

	return createdEmployee, nil
}

func UpdateEmployee(employee models.UpdateEmployeeDTO) (models.Employee, error) {
	db := database.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return models.Employee{}, fmt.Errorf("failed to begin transaction: %v", err)
	}

	sqlUpdate := `UPDATE "Employees" SET phone = $2 WHERE id = $1;`

	result, err := db.Exec(sqlUpdate, employee.ID, employee.Phone)

	if err != nil {
		return models.Employee{}, fmt.Errorf("unable to execute the query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Employee{}, fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return models.Employee{}, fmt.Errorf("no employee found with id: %d", employee.ID)
	}

	var updatedEmployee models.Employee
	sqlSelect := `SELECT id, created_at, name, type, phone, "isWorking" FROM "Employees" WHERE id = $1`
	err = tx.QueryRow(sqlSelect, employee.ID).Scan(
		&updatedEmployee.ID,
		&updatedEmployee.CreatedAt,
		&updatedEmployee.Name,
		&updatedEmployee.Type,
		&updatedEmployee.Phone,
		&updatedEmployee.IsWorking,
	)
	if err != nil {
		return models.Employee{}, fmt.Errorf("unable to fetch updated employee: %v", err)
	}
	if err = tx.Commit(); err != nil {
		return models.Employee{}, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return updatedEmployee, nil
}

func DeleteEmployee(id int) error {
	db := database.GetDB()

	sqlStatement := `UPDATE "Employees" SET "isWorking" = false WHERE id = $1;`

	result, err := db.Exec(sqlStatement, id)
	
	if err != nil {
		return fmt.Errorf("unable to execute the delete query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with id: %d", id)
	}

	return nil
}
