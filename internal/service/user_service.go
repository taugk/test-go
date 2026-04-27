package service

import (
	"errors"
	"test-oldo/internal/model"
	"test-oldo/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *model.User) error {
	if user.Name == "" || user.Phone == "" {
		return errors.New("name dan phone tidak boleh kosong")
	}
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetByID(id uint) (model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Update(id uint, input *model.User) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Phone = input.Phone

	err = s.repo.Update(&user)
	return user, err
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
