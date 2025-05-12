package utils

import (
	"github.com/wunnaaung-dev/payroll-bre/models"
	"regexp"
)

func ValidateEmployeeData(employee models.CreateEmployeeDTO) []string {
	var errors []string

	errors = append(errors, validateName(employee.Name)...)
	errors = append(errors, validatePhone(employee.Phone)...)

	if employee.Type == "" {
		errors = append(errors, "Employee Type is required")
	}

	return errors
}

func CheckEmployeePhone(employee models.UpdateEmployeeDTO) []string {
	return validatePhone(employee.Phone)
}

func CheckTeacherData(teacher models.CreateTeacherDTO) [] string {
	var errors []string

	if teacher.Teacher_ID == 0 {
		errors = append(errors, "Teacher ID is required")
	}

	if teacher.Subject == "" {
		errors = append(errors, "Subject is required")
	}

	if teacher.Role == "" {
		errors = append(errors, "Role is required")
	}

	if teacher.Total_Classes_Per_Month == 0 {
		errors = append(errors, "Total calls is required")
	}

	return errors
}

func validateName(name string) []string {
	var errors []string
	if name == "" {
		errors = append(errors, "Name is required")
	}
	return errors
}

func validatePhone(phone string) []string {
	var errors []string
	if phone == "" {
		errors = append(errors, "Phone is required")
	} else if !isValidPhoneNumber(phone) {
		errors = append(errors, "Phone number format is invalid")
	}
	return errors
}

func isValidPhoneNumber(phone string) bool {
	// Example regex for phone number validation
	phoneRegex := `^(\+)?[0-9]{10,15}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
