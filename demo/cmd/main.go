package main

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	pt "github.com/ptechen/kratos-proto/demo/api"
	"kratos-client-demo/demo/client"
)

//func main()  {
//	cnf := &warden.ClientConfig{}
//	c, err := client.NewClient(cnf)
//	if err != nil {
//		fmt.Println("1")
//		panic(err)
//	}
//	if c == nil {
//		fmt.Println("2")
//		panic(err)
//	}
//	req := &pb.HelloReq{Name: "321"}
//	res, err := c.SayHello(context.Background(), req)
//	if err != nil {
//		fmt.Println("3")
//		panic(err)
//	}
//	fmt.Println(res)
//	ress, err := c.HelloWorld(context.Background(), req)
//	if err != nil {
//		fmt.Println("4")
//		panic(err)
//	}
//	fmt.Println(ress)
//}

func main() {
	//config := &clientv3.Config{
	//	Endpoints:   []string{"192.168.3.241:2379"},
	//	DialTimeout: time.Second * 3,
	//	DialOptions: []grpc.DialOption{grpc.WithBlock()},
	//}
	//builder, err := etcd.New(config)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(builder)
	//resolver := builder.Build("demo.service")
	//cfg := &warden.ClientConfig{}
	//info, f := resolver.Fetch(context.TODO())
	//fmt.Println(info, f)
	cfg := &warden.ClientConfig{
		Dial:                   0,
		Timeout:                0,
		Breaker:                nil,
		Method:                 nil,
		Clusters:               nil,
		Zone:                   "",
		Subset:                 0,
		NonBlock:               false,
		KeepAliveInterval:      0,
		KeepAliveTimeout:       0,
		KeepAliveWithoutStream: false,
	}

	c, err := client.NewClient(cfg)
	if err != nil {
		return
	}
	he := &pt.HelloReq{Name: "123"}
	res, err := c.HelloWorld(context.Background(), he)
	if err != nil {
		return
	}
	fmt.Println(res)
}
