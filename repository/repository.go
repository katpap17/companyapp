package repository

import "gorm.io/gorm"

type DBHandler interface {
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
}

type Repository struct {
	db DBHandler
}
