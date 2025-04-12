package handlers

import "gorm.io/gorm"

type handlerWithDb struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handlerWithDb {
	return handlerWithDb{db}
}
