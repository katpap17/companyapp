package actions

import (
	"errors"

	"github.com/katpap17/companyapp/auth"
	"github.com/katpap17/companyapp/repository"
	"github.com/katpap17/companyapp/utils"
)

func Login(user *repository.User) (string, error) {
	var token string
	dbUser, err := repository.GetUser(user.Username)
	if err != nil {
		utils.Logger.Error("Failed to fetch user with error: ", err)
		return "", err
	}
	if dbUser == nil {
		utils.Logger.Error("User not found")
		return "", nil
	}
	if utils.ComparePasswords(dbUser.Password, user.Password) {
		token, err = auth.GenerateToken(user.Username)
		if err != nil {
			utils.Logger.Error("Failed to generate token with error: ", err)
			return "", err
		}
		utils.Logger.Info("Token generated successfully for user ", user.Username)
		return token, nil
	}
	utils.Logger.Info("Token not generated for user ", user.Username)
	return token, nil
}

func CreateUser(username string, password string) error {
	user, _ := repository.GetUser(username)
	if user != nil {
		return errors.New("User already exists")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		utils.Logger.Error("Password can not be hashed ", err)
		return err
	}
	err = repository.CreateUser(username, hashedPassword)
	utils.Logger.Info("User created succesfully ", username)
	return nil
}
