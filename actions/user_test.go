package actions

import (
	"testing"

	"github.com/katpap17/companyapp/repository"
	"github.com/katpap17/companyapp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestLogin_WrongPass(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetUserRepository(mockDB)
	mockUser := &repository.User{Username: "username", Password: "$2a$10$s.NE7lLIfw1qB6XKULPE9.Zi2k5Rt8wsY43cFyuYVBZXtuG6WWHXe"}
	mockDB.On("First", mock.AnythingOfType("*repository.User"), mock.Anything).Run(
		func(args mock.Arguments) {
			userArg := args.Get(0).(*repository.User)
			*userArg = *mockUser
		},
	).Return(&gorm.DB{})
	user := &repository.User{Username: "username", Password: "12346"}

	// Act
	token, err := Login(user)

	// Assert
	assert.Empty(t, token)
	assert.Nil(t, err)
}

func TestLogin_NoUser(t *testing.T) {
	// Arrange
	mockDB := new(utils.MockDB)
	repository.SetUserRepository(mockDB)
	mockDB.On("First", mock.AnythingOfType("*repository.User"), mock.Anything).Return(&gorm.DB{})
	user := &repository.User{Username: "username", Password: "12346"}

	// Act
	token, err := Login(user)

	// Assert
	assert.Empty(t, token)
	assert.Nil(t, err)
}
