module example.com/grpcdemo

go 1.13

replace google.golang.org/grpc v1.33.0 => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	golang.org/x/net v0.0.0-20210421230115-4e50805a0758 // indirect
	google.golang.org/grpc v1.33.0
)
