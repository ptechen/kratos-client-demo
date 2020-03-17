package client

import (
	"context"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	pb "github.com/ptechen/kratos-proto/demo/api"
	"google.golang.org/grpc"
	"sync"
)

// AppID .
var clientAppID = "direct://default/127.0.0.1:9000"

func SetAppId(appID string)  {
	once := sync.Once{}
	once.Do(func() {
		if appID != "" {
			clientAppID = appID
		}
	})
}

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (pb.DemoClient, error) {
	//opts = append(opts, grpc.WithTransportCredentials(addTlsClient()))
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), clientAppID)
	if err != nil {
		return nil, err
	}
	return pb.NewDemoClient(cc), nil
}
