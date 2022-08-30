package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"goTestProject/consulDemo/base"
	trippb "goTestProject/grpc-demo/gen/go"
	trip "goTestProject/grpc-demo/tripService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

func main() {
	go startGRPCGateway()
	// 服务发现
	address, port := base.GetServicesWithFilter(`Service=="darr_en1"`)

	fmt.Println(address, port)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	// 为grpc server 设置拦截器
	s := grpc.NewServer(
	//grpc.UnaryInterceptor( //为grpc server 设置拦截器
	//	//为grpc server 设置拦截器
	//	func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//		return nil, nil
	//	},
	//),
	)
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	// 为grpc 服务注册 HealthCheck 接口
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	// 服务注册  address 拿本机ip(其他机器能通过这个访问到你)
	base.RegisterWithGRPCHealthCheck("172.25.40.121", 8081, "grpcServer", []string{"grpc", "test"}, "grpcServer")

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGRPCGateway() {
	c := context.Background()
	// 创建可以cancel 的 上下文
	c, cancel := context.WithCancel(c)
	defer cancel()

	// 分发器，将http请求分发到grpc server
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		// protoc 序列化枚举类型 使用数字 1 not NOT_STARTED
		&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{UseEnumNumbers: true}},
	),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8081", opts)
	if err != nil {
		log.Fatalf("cannot grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot Listen and server: %v", err)
	}
}
