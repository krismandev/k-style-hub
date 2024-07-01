package repository

import (
	"k-style-test/model"
	"k-style-test/utility"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

// NewAuthRepository function to create a new instance of AuthRepositoryImpl
func NewAuthRepository(db *gorm.DB) AuthLogRepository {
	return &AuthRepositoryImpl{DB: db}
}

// CreateAuthLog method implementation
func (r *AuthRepositoryImpl) CreateAuthLog(authLog *model.AuthLog) error {
	var err error
	result := r.DB.Create(&authLog)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}

	return err
}

func (r *AuthRepositoryImpl) GetAuthLog(authLogs *[]model.AuthLog, params map[string]interface{}) error {
	var err error

	result := r.DB.Scopes(utility.Paginate(params)).Preload("User").Scopes(utility.Order(params)).Find(&authLogs)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err
}

func (r *AuthRepositoryImpl) AuthLogCount(count *int64) error {
	var err error

	result := r.DB.Model(&model.AuthLog{}).Count(count)
	err = utility.CheckErrorResult(result)
	if err != nil {
		logrus.Errorf("Error in Repository: %v", err)
	}
	return err

}
