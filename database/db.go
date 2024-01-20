package database

import (
	"github.com/katpap17/companyapp/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres password=mysecretpassword dbname=companydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	return db, nil
}
