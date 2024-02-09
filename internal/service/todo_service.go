package service

import (
	"start/internal/model"
	"start/internal/repository"
)

type TodoService struct {
	repo repository.TodoUsersRepository
}

func NewTodoService(repo repository.TodoUsersRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s TodoService) SetUserService(user *model.UserTodo) error {
	err := s.repo.SetUserStorage(user)
	if err != nil {
		return ErrorAddUser
	}
	return nil
}

func (s TodoService) GetUserByIDService(id int) (*model.UserTodo, error) {
	user, err := s.repo.GetUserByIDStorage(id)
	if err != nil {
		return nil, ErrorReceivingUser
	}
	return user, nil
}

func (s TodoService) CheckUserByIDService(id int) (bool, error) {
	exists, err := s.repo.CheckUserByIDStorage(id)
	if err != nil {
		return false, ErrorCheckUser
	}
	return exists, nil
}

func (s TodoService) DeleteUserByIdService(id int) error {
	exists, err := s.repo.CheckUserByIDStorage(id)
	if err != nil {
		return ErrorCheckUser
	}
	if !exists {
		return ErrorNotFoundUser
	}
	err = s.repo.DeleteUserByIdStorage(id)
	if err != nil {
		return ErrorDeleteUser
	}
	return nil
}

func (s TodoService) GetAllUserIDService() ([]int, error) {
	ids, err := s.repo.GetAllUserIDStorage()
	if err != nil {
		return nil, ErrorReceivingUser
	}
	return ids, nil
}
