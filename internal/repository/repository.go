package repository

import (
	"github.com/jmoiron/sqlx"
	"start/internal/model"
	"start/internal/repository/postgres"
)

type AuthorizationRepository interface {
	CreateUserStorageAuth(user model.UserAuth) error
	GetUserStorageAuth(username, password, email string) (model.UserAuth, error)
	CheckUserByUsernameAndPassword(username, password string) (bool, error)
	CheckUserByUsername(username string) (bool, error)
	CheckUserByEmail(email string) (bool, error)
}

type TodoUsersRepository interface {
	SetUserStorage(user *model.UserTodo) error
	GetUserByIDStorage(id int) (*model.UserTodo, error)
	CheckUserByIDStorage(id int) (bool, error)
	DeleteUserByIdStorage(id int) error
	GetAllUserIDStorage() ([]int, error)
}

type Repository struct {
	Auth AuthorizationRepository
	Todo TodoUsersRepository
}

func NewStorageUserPostgres(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(db),
		Todo: postgres.NewTodoPostgres(db),
	}
}
