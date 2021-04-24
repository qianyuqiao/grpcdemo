package common

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"time"
)

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	// 如果没有设置超时 自动加上超时
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
	}

	p := peer.Peer{}
	if opts == nil {
		opts = []grpc.CallOption{grpc.Peer(&p)}
	} else {
		opts = append(opts, grpc.Peer(&p))
	}

	start := time.Now()
	defer func() {
		in, _ := json.Marshal(req)
		out, _ := json.Marshal(reply)
		inStr, outStr := string(in), string(out)
		duration := int64(time.Since(start) / time.Millisecond)

		var remoteServer string
		if p.Addr != nil {
			remoteServer=p.Addr.String()
		}

		log.Println("grpc", method, "in", inStr, "out", outStr, "err", err, "duration/ms", duration, "remote_server", remoteServer)

	}()

	return invoker(ctx, method, req, reply, cc, opts...)
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	log.Println("ip", remoteAddr, "grpc_access_start", info.FullMethod, "in", inStr)

	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			log.Println("ip", remoteAddr, "grpc_access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		} else {
			log.Println("ip", remoteAddr, "grpc_access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		}
	}()

	resp, err = handler(ctx, req)

	return
}
