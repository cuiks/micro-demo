package main

import (
	"github.com/cuiks/user/domain/repository"
	service2 "github.com/cuiks/user/domain/service"
	"github.com/cuiks/user/handler"
	pb "github.com/cuiks/user/proto"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	service = "micro.service.user"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	db, err := gorm.Open("mysql", "root:aaa123456@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)

	repo := repository.NewUserRepo(db)
	//repo.InitTable()

	userService := service2.NewUserService(repo)
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{Service: userService})
	if err != nil {
		panic(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
