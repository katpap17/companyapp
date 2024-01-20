package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Arrange
	userName := "username"

	// Act
	token, err := GenerateToken(userName)

	// Assert
	assert.Empty(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateJWT_ValidToken(t *testing.T) {
	// Arrange
	userID := "username"
	token, _ := GenerateToken(userID)

	// Act
	claims, valid := ValidateToken(token)

	// Assert
	assert.NotEmpty(t, claims)
	assert.Equal(t, userID, *&claims.Username)
	assert.Equal(t, true, valid)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	// Arrange
	token := ""

	// Act
	resultUserID, valid := ValidateToken(token)

	// Assert
	assert.Empty(t, resultUserID)
	assert.Equal(t, false, valid)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	// Arrange
	userID := "username"
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(-time.Hour).Unix(),
		Id:        userID,
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := expiredToken.SignedString(jwtKey)

	// Act
	resultUserID, err := ValidateToken(signedToken)

	// Assert
	assert.NotNil(t, err)
	assert.Empty(t, resultUserID)
}
