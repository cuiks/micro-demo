package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	"newmicro/proto/kshcui"
)

func main() {
	// 实例化
	service := micro.NewService(
		micro.Name("kshcui.hello.client"),
	)
	// 初始化
	service.Init()

	helloService := kshcui.NewHelloService("kshcui.hello.server", service.Client())
	res, err := helloService.SayHello(context.TODO(), &kshcui.SayRequest{Msg: "你好"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
