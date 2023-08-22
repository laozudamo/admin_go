package main

/*
1.取出server
2.挂载方法
3.注册服务
4.创建监听
*/
import (
	hello_rpc "admin_go/grpc_class/pb"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	hello_rpc.UnimplementedHelloGrpcServer
}

func main() {
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	s := grpc.NewServer()
	hello_rpc.RegisterHelloGrpcServer(s, &Server{})

	s.Serve(l)
}

func (s *Server) SayHi(ctx context.Context, req *hello_rpc.Req) (res *hello_rpc.Res, err error) {
	fmt.Println("req:", req.GetMsg())
	return &hello_rpc.Res{Msg: "服务端返回RPC数据"}, nil
}
