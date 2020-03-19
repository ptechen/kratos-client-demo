package client

import (
	"context"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/net/rpc/warden/resolver"
	"github.com/ptechen/kratos-proto/demo/api"
	"google.golang.org/grpc"
	"os"
)

// AppID your appid, ensure unique.
const AppID = "demo.service" // NOTE: example

func init() {
	// NOTE: 注意这段代码，表示要使用discovery进行服务发现
	// NOTE: 还需注意的是，resolver.Register是全局生效的，所以建议该代码放在进程初始化的时候执行
	// NOTE: ！！！切记不要在一个进程内进行多个不同中间件的Register！！！
	// NOTE: 在启动应用时，可以通过flag(-discovery.nodes) 或者 环境配置(DISCOVERY_NODES)指定discovery节点
	resolver.Register(discovery.Builder())
}

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (api.DemoClient, error) {
	client := warden.NewClient(cfg, opts...)
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	conn, err := client.DialTLS(context.Background(), "discovery://default/"+AppID, "./certs/server.pem", name)
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return api.NewDemoClient(conn), nil
}