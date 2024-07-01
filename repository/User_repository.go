package repository

import (
	model "k-style-test/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(user *[]model.User) error
	GetUser(user *model.User) error
	CreateUser(tx *gorm.DB, user *model.User) error
}
