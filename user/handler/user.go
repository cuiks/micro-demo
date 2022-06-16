package handler

import (
	context "context"
	"github.com/cuiks/user/domain/model"
	"github.com/cuiks/user/domain/service"
	pb "github.com/cuiks/user/proto"
)

type User struct {
	Service service.IUserService
}

func (u *User) UserRegister(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	user := &model.User{
		UserName: request.Username,
		Age:      request.Age,
	}
	uid, err := u.Service.AddUser(user)
	if err != nil {
		return err
	}
	response.Id = uid
	response.Username = user.UserName
	return nil
}

func (u *User) FindUser(ctx context.Context, request *pb.UserRequest, response *pb.UserResponse) error {
	user, err := u.Service.FindUserByName(request.User.Username)
	if err != nil {
		return err
	}
	response.Id = user.ID
	response.Age = user.Age
	response.Username = user.UserName
	return nil
}
