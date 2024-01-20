package repository

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/katpap17/companyapp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	SetCompanyRepository(mockDB)

	mockCompany := &Company{Description: "Description"}
	id, _ := uuid.NewV7()
	mockDB.On("First", mock.AnythingOfType("*repository.Company"), mock.Anything).Run(
		func(args mock.Arguments) {
			compArg := args.Get(0).(*Company)
			*compArg = *mockCompany
		},
	).Return(&gorm.DB{})
	conds := []interface{}{id}

	// Act
	company, err := GetCompany(id)

	// Assert
	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*repository.Company"), conds)
	assert.Nil(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, "Description", company.Description)

}

func TestCreateCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	SetCompanyRepository(mockDB)
	registered := true
	cType := Cooperative
	company := &Company{
		Name:        "Test Company",
		Description: "Test Description",
		Employees:   400,
		Registered:  &registered,
		CompanyType: &cType,
	}
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	// Act
	err := CreateCompany(company)

	// Assert
	mockDB.AssertCalled(t, "Create", company)
	assert.Nil(t, err)
}

func TestUpdateCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	SetCompanyRepository(mockDB)
	registered := true
	cType := Cooperative
	company := &Company{
		Name:        "Test Company",
		Description: "Test Description",
		Employees:   400,
		Registered:  &registered,
		CompanyType: &cType,
	}
	mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

	// Act
	err := UpdateCompany(company)

	// Assert
	mockDB.AssertCalled(t, "Save", company)
	assert.Nil(t, err)
}

func TestDeleteCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	SetCompanyRepository(mockDB)
	id, _ := uuid.NewV7()
	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{})
	conds := []interface{}{id}

	// Act
	err := DeleteCompany(id)

	// Assert
	mockDB.AssertCalled(t, "Delete", mock.AnythingOfType("*repository.Company"), conds)
	assert.Nil(t, err)

}
