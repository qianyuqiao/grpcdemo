package main

import (
	"example.com/grpcdemo/app"
	"example.com/grpcdemo/grpc/common"
	"example.com/grpcdemo/grpc/routeguide"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("p", 12305, "port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(common.UnaryServerInterceptor))
	routeguide.RegisterRouteGuideServer(s, app.NewServer())
	s.Serve(lis)
}
