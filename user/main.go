package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/cuiks/user/common"
	"github.com/cuiks/user/domain/repository"
	service2 "github.com/cuiks/user/domain/service"
	"github.com/cuiks/user/handler"
	pb "github.com/cuiks/user/proto"
	opentracing2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	service = "micro.service.user"
	version = "latest"
)

func main() {
	// 配置中心
	config, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// 链路追踪
	trace, io, err := common.NewTrace("micro.service.user", "127.0.0.1:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(trace)

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		// 地址和端口
		micro.Address("127.0.0.1:8002"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()

	// 获取配置
	mysqlInfo := common.GetMysqlConfig(config, "mysql")

	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlInfo.User,
			mysqlInfo.Pwd,
			mysqlInfo.Host,
			mysqlInfo.Port,
			mysqlInfo.Database,
		))
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
