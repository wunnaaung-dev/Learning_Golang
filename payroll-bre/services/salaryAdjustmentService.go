package services

import (
	"fmt"
	"github.com/wunnaaung-dev/payroll-bre/database"
	"github.com/wunnaaung-dev/payroll-bre/models"

	"github.com/hyperjumptech/grule-rule-engine/ast"

	"github.com/hyperjumptech/grule-rule-engine/builder"

	"github.com/hyperjumptech/grule-rule-engine/engine"

	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func GetSalaryAdjustment(id int) (models.SalaryAdjustment, error) {
	db := database.GetDB()

	sqlStatement := `
		SELECT 
			"Payroll_Adjustments".employee_id,
			"Employees".name,
			"Employees".type,
			"Payroll_Adjustments".records,
			"Payroll_Adjustments".created_at
		FROM 
			"Employees"
		INNER JOIN 
			"Payroll_Adjustments"
			ON "Employees".id = "Payroll_Adjustments".employee_id
		WHERE
			"Employees".id = $1
			AND DATE_TRUNC('month', "Payroll_Adjustments".created_at) = DATE_TRUNC('month', CURRENT_DATE);
	`

	var adjustment models.SalaryAdjustment
	err := db.QueryRow(sqlStatement, id).Scan(
		&adjustment.Employee_ID,
		&adjustment.Name,
		&adjustment.Type,
		&adjustment.Records,
		&adjustment.CreatedAt,
	)
	if err != nil {
		return models.SalaryAdjustment{}, fmt.Errorf("unable to fetch adjustment info: %v", err)
	}

	return adjustment, nil
}

func TestingRule(employeeId int) string {
	empInfo, err := GetEmployeeInfo(employeeId)
	if err != nil {
		return fmt.Sprintf("Error getting employee info: %v", err)
	}

	result := &models.Result{}

	dataCtx := ast.NewDataContext()
	err = dataCtx.Add("Employee", &empInfo)
	if err != nil {
		return fmt.Sprintf("Error adding employee to data context: %v", err)
	}
	err = dataCtx.Add("Result", result)
	if err != nil {
		return fmt.Sprintf("Error adding result to data context: %v", err)
	}

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)

	fileRes := pkg.NewFileResource("rules/bonus_check.grl")
	err = rb.BuildRuleFromResource("BonusRules", "0.0.1", fileRes)
	if err != nil {
		return fmt.Sprintf("Error building rules: %v", err)
	}

	kb, _ := lib.NewKnowledgeBaseInstance("BonusRules", "0.0.1")
	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, kb)

	if err != nil {
		return fmt.Sprintf("Error executing rules: %v", err)
	}

	// 6. Return the final result (e.g., a message or value set by rules)
	return result.Message

}

func CreateSalaryAdjustment(dto models.CreateAdjustmentDTO) (models.SalaryAdjustment, error) {
	db := database.GetDB()

	sqlStatement := `
		INSERT INTO "Payroll_Adjustments" (employee_id, records)
		VALUES ($1, $2)
	`

	_, err := db.Exec(sqlStatement, dto.Employee_ID, dto.Records)
	if err != nil {
		return models.SalaryAdjustment{}, fmt.Errorf("unable to create salary adjustment: %v", err)
	}

	adjustment, err := GetSalaryAdjustment(int(dto.Employee_ID))
	if err != nil {
		return models.SalaryAdjustment{}, err
	}

	return adjustment, nil
}

func UpdateSalaryAdjustment(id int, dto models.UpdateAdjustmentDTO) (models.SalaryAdjustment, error) {
	db := database.GetDB()

	sqlStatement := `
		UPDATE "Payroll_Adjustments"
		SET records = $1
		WHERE employee_id = $2
		AND DATE_TRUNC('month', created_at) = DATE_TRUNC('month', CURRENT_DATE)
	`

	_, err := db.Exec(sqlStatement, dto.Records, id)
	if err != nil {
		return models.SalaryAdjustment{}, fmt.Errorf("unable to update salary adjustment: %v", err)
	}

	adjustment, err := GetSalaryAdjustment(id)
	if err != nil {
		return models.SalaryAdjustment{}, err
	}

	return adjustment, nil
}
