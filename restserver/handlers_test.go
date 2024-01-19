package restserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCompany(t *testing.T) {
	// Arrange
	request, _ := http.NewRequest("GET", "/companies/5", nil)
	responseRecorder := httptest.NewRecorder()
	expectedResult := `{"company":{"ID":"00000000-0000-0000-0000-000000000000","Name":"XM","Description":"5","Employees":0,"Registered":false,"Type":0},"status":"success"}
`

	// Act
	GetCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
	assert.Equal(t, expectedResult, responseRecorder.Body.String())
}

func TestCreateCompany(t *testing.T) {
	// Arrange
	body := `{"company":{"Name":"","Description":"5","Employees":0,"Registered":false,"Type":0}}`
	request, _ := http.NewRequest("POST", "/companies", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()

	// Act
	CreateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}

func TestUpdateCompany(t *testing.T) {
	// Arrange
	body := `{"company":{"Name":"","Description":"5","Employees":0,"Registered":false,"Type":0},"status":"success"}`
	request, _ := http.NewRequest("PATCH", "/companies/00000000-0000-0000-0000-000000000000", bytes.NewBufferString(body))
	responseRecorder := httptest.NewRecorder()

	// Act
	UpdateCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}

func TestDeleteCompany(t *testing.T) {
	// Arrange
	request, _ := http.NewRequest("DELETE", "/companies/00000000-0000-0000-0000-000000000000", nil)
	responseRecorder := httptest.NewRecorder()

	// Act
	DeleteCompany(responseRecorder, request)

	// Assert
	assert.Equal(t, 200, responseRecorder.Result().StatusCode)
}
