package main

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	pb "github.com/ptechen/kratos-proto/demo/api"
	"kratos-client-demo/demo/client"
)

func main()  {
	cnf := &warden.ClientConfig{}
	c, err := client.NewClient(cnf)
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	if c == nil {
		fmt.Println("2")
		panic(err)
	}
	req := &pb.HelloReq{Name: "321"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println("3")
		panic(err)
	}
	fmt.Println(res)
	ress, err := c.HelloWorld(context.Background(), req)
	if err != nil {
		fmt.Println("4")
		panic(err)
	}
	fmt.Println(ress)
}