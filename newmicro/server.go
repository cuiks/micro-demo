package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	hello "newmicro/proto/kshcui"
)

type HelloServer struct {
}

// SayHello 需要实现的方法
func (h *HelloServer) SayHello(ctx context.Context, req *hello.SayRequest, rsp *hello.SayResponse) error {
	rsp.Answer = "发送成功: " + req.Msg
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("kshcui.hello.server"),
	)
	// 初始化方法
	service.Init()
	// 注册服务
	hello.RegisterHelloHandler(service.Server(), new(HelloServer))
	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
