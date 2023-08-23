package main

import (
	hello_rpc "admin_go/grpc_class/pb"
	"context"

	"google.golang.org/grpc"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("localhost:8080", "")

	//if err != nil {
	//	log.Fatalf("Failed to load TLS certificates: %v", err)
	//}
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := hello_rpc.NewHelloGrpcClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_rpc.Req{Msg: "客户端发送"})
	println(req.GetMsg())
}
