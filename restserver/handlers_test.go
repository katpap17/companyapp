package restserver

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/katpap17/companyapp/auth"
	"github.com/katpap17/companyapp/repository"
	"github.com/katpap17/companyapp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	id, _ := uuid.NewV7()
	registered := false
	companyType := repository.Cooperative
	mockCompany := &repository.Company{ID: id, Name: "XM", Description: "Description", Employees: 100, Registered: &registered, CompanyType: &companyType}
	expectedResult := fmt.Sprintf(`{"company":{"id":"%s","name":"XM","description":"Description","employees":100,"registered":false,"companyType":2},"status":"success"}
`, id)
	mockDB.On("First", mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			compArg := args.Get(0).(*repository.Company)
			*compArg = *mockCompany
		},
	).Return(&gorm.DB{})
	request, _ := http.NewRequest("GET", "/companies/"+id.String(), nil)
	responseRecorder := httptest.NewRecorder()

	// Act
	GetCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
	assert.Equal(t, expectedResult, responseRecorder.Body.String())
}

func TestCreateCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	body := `{"name":"XM","description":"5","employees":100,"registered":false,"companyType":0}`
	request, _ := http.NewRequest("POST", "/companies", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	// Act
	CreateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}

func TestCreateCompany_FailValidation(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	body := `{"name":"XM","description":"5","registered":false,"companyType":0}`
	request, _ := http.NewRequest("POST", "/companies", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	// Act
	CreateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 400, responseRecorder.Result().StatusCode)
}

func TestUpdateCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	body := `{"name":"XM","description":"5","employees":100,"registered":false,"companyType":0}`
	request, _ := http.NewRequest("PATCH", "/companies/a1f5e7ab-8b2a-4f48-bab5-9de29c2638a2", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()
	mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

	// Act
	UpdateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}

func TestUpdateCompany_FailValidation(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	body := `{"name":"XM","description":"5","employees":100,"companyType":0}`
	request, _ := http.NewRequest("PATCH", "/companies/a1f5e7ab-8b2a-4f48-bab5-9de29c2638a2", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()
	mockDB.On("Save", mock.Anything).Return(&gorm.DB{})

	// Act
	UpdateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 400, responseRecorder.Result().StatusCode)
}

func TestDeleteCompany(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetCompanyRepository(mockDB)
	request, _ := http.NewRequest("DELETE", "/companies/a1f5e7ab-8b2a-4f48-bab5-9de29c2638a2", nil)
	responseRecorder := httptest.NewRecorder()
	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{})

	// Act
	DeleteCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}

func TestLogin(t *testing.T) {
	// Arrange
	responseRecorder := httptest.NewRecorder()
	mockDB := new(utils.MockDB)
	repository.SetUserRepository(mockDB)
	mockUser := &repository.User{Username: "username", Password: "$2a$10$s.NE7lLIfw1qB6XKULPE9.Zi2k5Rt8wsY43cFyuYVBZXtuG6WWHXe"}
	mockDB.On("First", mock.AnythingOfType("*repository.User"), mock.Anything).Run(
		func(args mock.Arguments) {
			userArg := args.Get(0).(*repository.User)
			*userArg = *mockUser
		},
	).Return(&gorm.DB{})
	request, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(`{"username":"katerina","password":"12345"}`))
	token, _ := auth.GenerateToken("username")
	request.Header.Set("Authorization", "Bearer "+token)

	// Act
	Login(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}
