package routeguide

import (
	"context"
	"example.com/grpcdemo/config"
	"example.com/grpcdemo/grpc/common"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var client routeGuideClient

func init() {
	var err error
	client.cc, err = grpc.Dial(
		"127.0.0.1:12305",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(common.UnaryClientInterceptor)),
	)
	if err != nil {
		panic(err)
	}

	config.Mode = config.ModeQa

}

func Ping(ctx context.Context, in *PingRequest) (*PingReply, error) {
	out := new(PingReply)
	err := client.cc.Invoke(ctx, "/routeguide.RouteGuide/Ping", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
