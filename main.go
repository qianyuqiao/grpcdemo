package main

import (
	"example.com/grpcdemo/app"
	"example.com/grpcdemo/config"
	"example.com/grpcdemo/grpc/common"
	"example.com/grpcdemo/grpc/routeguide"
	"example.com/grpcdemo/utils/dlog"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port = flag.Int("p", 12305, "port")
)

func main() {
	topic := "grpcdemo" // set log topic to project or service name
	dlog.SetTopic(topic)
	dlog.DebugLog(true)

	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	// bind OS events to the signal channel
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(common.UnaryServerInterceptor))
	routeguide.RegisterRouteGuideServer(s, app.NewServer())
	go func() {
		// start serve
		if err := s.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	// block until either OS signal, or server fatal error
	select {
	case err := <-errChan:
		log.Printf("Fatal error: %v\n", err)
	case <-stopChan:
		s.GracefulStop()
		dlog.Close()
	}

}

func init() {
	config.Mode = config.ModeQa
}
