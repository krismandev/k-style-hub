package usecase

import (
	"context"
	"k-style-test/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	BeginTx(ctx context.Context, fn func(tx UnitOfWork) error) error
	NewUserRepository() repository.UserRepository
	NewCustomerRepository() repository.CustomerRepository
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (s *unitOfWork) BeginTx(ctx context.Context, fn func(tx UnitOfWork) error) error {
	var err error
	tx := s.db.Begin()
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback().Error; rbErr != nil {
				logrus.Errorf("tx err: %v, rb err: %v", err, rbErr)
				logrus.Infof("%v\n", err)
			}
		} else {
			err = tx.Commit().Error
			logrus.Printf("%v\n", err)
		}
	}()

	newUnitOfWork := &unitOfWork{
		db: tx,
	}
	err = fn(newUnitOfWork)
	return nil
}

func (s *unitOfWork) NewUserRepository() repository.UserRepository {
	return &repository.UserRepositoryImpl{
		DB: s.db,
	}
}

func (s *unitOfWork) NewCustomerRepository() repository.CustomerRepository {
	return &repository.CustomerRepositoryImpl{
		DB: s.db,
	}
}
