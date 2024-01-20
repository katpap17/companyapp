package repository

import (
	"github.com/katpap17/companyapp/utils"
)

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user_table"
}

type UserRepository struct {
	Repository Repository
}

var userRepository UserRepository

func SetUserRepository(db DBHandler) {
	userRepository = UserRepository{Repository: Repository{db: db}}
}

func GetUser(username string) (*User, error) {
	return userRepository.get(username)
}

func CreateUser(username string, password string) error {
	return userRepository.create(username, password)
}

func (r *UserRepository) get(username string) (*User, error) {
	var user User
	if err := r.Repository.db.First(&user, "username = ?", username).Error; err != nil {
		utils.Logger.Error(err.Error())
		return nil, err
	}
	return &user, nil

}

func (r *UserRepository) create(username string, password string) error {
	user := User{Username: username, Password: password}
	result := r.Repository.db.Create(user)
	if result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil

}
