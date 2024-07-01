package repository

import "k-style-test/model"

type AuthLogRepository interface {
	CreateAuthLog(authLog *model.AuthLog) error
	GetAuthLog(authLogs *[]model.AuthLog, params map[string]interface{}) error
	AuthLogCount(result *int64) error
}
