package service

import (
	"github.com/cuiks/user/domain/model"
	"github.com/cuiks/user/domain/repository"
)

type IUserService interface {
	AddUser(user *model.User) (int64, error)
	FindUserByName(string) (*model.User, error)
}

type UserService struct {
	repo repository.IUserRepo
}

// bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

func (u *UserService) AddUser(user *model.User) (int64, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) FindUserByName(s string) (*model.User, error) {
	return u.repo.FindUserByName(s)
}

func NewUserService(repo repository.IUserRepo) IUserService {
	return &UserService{repo: repo}
}
