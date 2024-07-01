package utility

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CheckErrorResult(result *gorm.DB) error {
	var err error
	err = result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &NotFoundError{Code: 404, Message: "Data not Found"}
		} else {
			return err
		}
	}
	return err
}

func RollbackInUseCase(tx *gorm.DB, err error) {
	logrus.Info("Error in UseCase : %v Rolling back transaction", err)
	tx.Rollback()
}

func NullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func returnEmptyIfNil(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
