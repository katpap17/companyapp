package restserver

import (
	"testing"

	"github.com/katpap17/companyapp/repository"
	"github.com/stretchr/testify/assert"
)

func TestValidateCompany(t *testing.T) {
	// Arrange
	registered := true
	ctype := repository.Cooperative
	tests := []struct {
		name          string
		inputCompany  repository.Company
		expectedValid bool
	}{
		{
			name: "Valid Company",
			inputCompany: repository.Company{
				Name:        "Example Company",
				Description: "A description",
				Employees:   10,
				Registered:  &registered,
				CompanyType: &ctype,
			},
			expectedValid: true,
		},
		{
			name: "Invalid Company",
			inputCompany: repository.Company{
				Name:        "Example Company",
				Description: "A description",
				Registered:  &registered,
				CompanyType: &ctype,
			},
			expectedValid: false,
		},
		{
			name: "Invalid Company",
			inputCompany: repository.Company{
				Name:        "Example Company",
				Description: "A description",
				Employees:   10,
				CompanyType: &ctype,
			},
			expectedValid: false,
		},
		{
			name: "Invalid Company",
			inputCompany: repository.Company{
				Name:        "Example Company",
				Description: "A description",
				Employees:   10,
				Registered:  &registered,
			},
			expectedValid: false,
		},
		{
			name: "Invalid Company",
			inputCompany: repository.Company{
				Description: "A description",
				Employees:   10,
				Registered:  &registered,
				CompanyType: &ctype,
			},
			expectedValid: false,
		},
		{
			name: "Invalid Company",
			inputCompany: repository.Company{
				Name:        "Name that is too long",
				Description: "A description",
				Employees:   10,
				Registered:  &registered,
				CompanyType: &ctype,
			},
			expectedValid: false,
		},
	}
	// Act Assert
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateCompany(tt.inputCompany)
			assert.Equal(t, tt.expectedValid, result)
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name          string
		inputUser     repository.User
		expectedValid bool
	}{
		{
			name: "Valid User",
			inputUser: repository.User{
				Username: "john_doe",
				Password: "secure_password",
			},
			expectedValid: true,
		},
		{
			name: "Invalid Username",
			inputUser: repository.User{
				Username: "",
				Password: "secure_password",
			},
			expectedValid: false,
		},
		{
			name: "Invalid Password",
			inputUser: repository.User{
				Username: "john_doe",
				Password: "",
			},
			expectedValid: false,
		},
	}
	// Act Assert
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateUser(tt.inputUser)
			assert.Equal(t, tt.expectedValid, result)
		})
	}
}
