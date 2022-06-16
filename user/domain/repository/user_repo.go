package repository

import (
	"github.com/cuiks/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepo interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	CreateUser(*model.User) (int64, error)
}

type UserRepo struct {
	mysqlDB *gorm.DB
}

func (u *UserRepo) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

func (u *UserRepo) FindUserByName(username string) (*model.User, error) {
	user := &model.User{}
	err := u.mysqlDB.Where("username=?", username).Find(user).Error
	return user, err
}

func (u UserRepo) CreateUser(user *model.User) (int64, error) {
	err := u.mysqlDB.Create(&user).Error
	return user.ID, err
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{mysqlDB: db}
}
