package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	// Arrange
	tests := []struct {
		name          string
		inputPassword string
		expectedError error
	}{
		{
			name:          "Valid Password",
			inputPassword: "secure_password",
			expectedError: nil,
		},
		{
			name:          "Empty Password",
			inputPassword: "The curious cat chased a fluttering butterfly through the enchanting garden, leaping gracefully over the vibrant flowers in pursuit of its whimsical prey.",
			expectedError: bcrypt.ErrPasswordTooLong,
		},
	}
	// Act Assert
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.inputPassword)
			assert.Equal(t, tt.expectedError, err)
			if err != tt.expectedError {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

func TestComparePasswords(t *testing.T) {
	// Arrange
	validPassword := "secure_password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
	invalidPassword := "wrong_password"

	tests := []struct {
		name          string
		encryptedPass string
		plainPass     string
		expectedMatch bool
	}{
		{
			name:          "Matching Passwords",
			encryptedPass: string(hashedPassword),
			plainPass:     validPassword,
			expectedMatch: true,
		},
		{
			name:          "Non-Matching Passwords",
			encryptedPass: string(hashedPassword),
			plainPass:     invalidPassword,
			expectedMatch: false,
		},
	}
	// Act Assert
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ComparePasswords(tt.encryptedPass, tt.plainPass)
			assert.Equal(t, tt.expectedMatch, result)
		})
	}

}
