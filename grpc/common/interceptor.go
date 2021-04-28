package common

import (
	"context"
	"encoding/json"
	"example.com/grpcdemo/utils/dlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
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
			// 通过peer在拦截器里拿到请求对应的远端服务器的IP和端口号
			// 对通过服务发现连接服务器的客户端的调试和记录非常有用。
			remoteServer=p.Addr.String()
		}

		dlog.Info("grpc", method, "in", inStr, "out", outStr, "err", err, "duration/ms", duration, "remote_server", remoteServer)

	}()

	return invoker(ctx, method, req, reply, cc, opts...)
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String() // 拿到远端发起请求的客户端的IP和端口号

	in, _ := json.Marshal(req)
	inStr := string(in)
	dlog.Info("ip", remoteAddr, "grpc_access_start", info.FullMethod, "in", inStr)

	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			dlog.Info("ip", remoteAddr, "grpc_access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		} else {
			dlog.Info("ip", remoteAddr, "grpc_access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
		}
	}()

	resp, err = handler(ctx, req)

	return
}
