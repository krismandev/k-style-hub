package repository

import (
	model "k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository UserRepositoryImpl) GetUser(user *model.User) error {
	var err error
	result := repository.DB.Where("email = ?", &user.Email).First(&user)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository UserRepositoryImpl) CreateUser(tx *gorm.DB, user *model.User) error {
	var err error
	result := tx.Create(&user)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (repository UserRepositoryImpl) GetUsers(users *[]model.User) error {
	var err error
	result := repository.DB.Find(&users)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}
