package app

import (
	"context"
	"example.com/grpcdemo/grpc/routeguide"
)

type server struct {

}

func (s server) Ping(ctx context.Context, request *routeguide.PingRequest) (reply *routeguide.PingReply, err error) {
	reply = &routeguide.PingReply{
		Reply:                "pong",
	}

	return
}

func NewServer() routeguide.RouteGuideServer  {
	return &server{}
}

