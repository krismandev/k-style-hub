package usecase

import (
	"context"
	"k-style-test/model/request"
	"k-style-test/model/response"
)

type AuthUseCase interface {
	Login(ctx context.Context, request request.LoginRequest) (response.LoginResponse, error)
	GetAuthLog(ctx context.Context, request request.AuthLogRequest) ([]response.AuthLogResponse, int64, error)
}
