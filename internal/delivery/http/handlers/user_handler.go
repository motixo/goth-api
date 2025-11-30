package handlers

import (
	"github.com/mot0x0/goth-api/internal/domain/service"
	"github.com/mot0x0/goth-api/internal/domain/usecase/user"
)

type UserHandler struct {
	usecase user.UseCase
	logger  service.Logger
}

func NewUserHandler(usecase user.UseCase, logger service.Logger) *UserHandler {
	return &UserHandler{
		usecase: usecase,
		logger:  logger,
	}
}
