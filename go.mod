module example.com/grpcdemo

go 1.13

replace google.golang.org/grpc v1.33.0 => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/gopherjs/gopherjs v0.0.0-20210420193930-a4630ec28c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jtolds/gls v4.20.0+incompatible
	github.com/labstack/gommon v0.3.0
	golang.org/x/net v0.0.0-20210421230115-4e50805a0758 // indirect
	google.golang.org/grpc v1.33.0
)
