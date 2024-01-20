package repository

import (
	"testing"

	"github.com/katpap17/companyapp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	SetUserRepository(mockDB)

	mockUser := &User{Username: "username"}
	mockDB.On("First", mock.AnythingOfType("*repository.User"), mock.Anything).Run(
		func(args mock.Arguments) {
			userArg := args.Get(0).(*User)
			*userArg = *mockUser
		},
	).Return(&gorm.DB{})

	// Act
	user, err := GetUser(mockUser.Username)

	// Assert
	mockDB.AssertCalled(t, "First", mock.AnythingOfType("*repository.User"), mock.Anything)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "username", user.Username)

}
