package service

import (
	"start/internal/model"
	"start/internal/repository"
)

type AuthorizationService interface {
	RegistrationService(user model.UserAuth) (string, error)
	LoginService(username, password string) (string, error)
}

type TodoUsersService interface {
	SetUserService(user *model.UserTodo) error
	GetUserByIDService(id int) (*model.UserTodo, error)
	CheckUserByIDService(id int) (bool, error)
	DeleteUserByIdService(id int) error
	GetAllUserIDService() ([]int, error)
}

type Service struct {
	Auth AuthorizationService
	Todo TodoUsersService
}

func NewServiceUser(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
		Todo: NewTodoService(repos.Todo),
	}
}
