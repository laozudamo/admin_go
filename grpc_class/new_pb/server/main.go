package main

import (
	"admin_go/grpc_class/new_pb/person"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

// PersonService 定义服务
type PersonService struct {
	person.UnimplementedSearchServiceServer
}

func (*PersonService) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	fmt.Println(req.GetName(), req.GetAge())
	return &person.PersonRes{
		Name: "服务端返回",
		Age:  20,
	}, nil
}

func (*PersonService) SearchIn(server person.SearchService_SearchInServer) error {
	// 接收客户端数据
	for {
		req, err := server.Recv()
		fmt.Println(req.GetName(), req.GetAge())
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "服务端接收数据完毕", Age: 0})
			break
		}
	}
	return nil
}
func (*PersonService) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	fmt.Println(req.GetName(), req.GetAge())
	i := 0

	for {
		if i > 10 {
			break
		}
		time.Sleep(time.Second)
		server.Send(&person.PersonRes{
			Name: "流发送",
			Age:  int32(i),
		})
		i++
	}
	return nil
}
func (*PersonService) SearchInOut(server person.SearchService_SearchInOutServer) error {
	var wg sync.WaitGroup // 定义一个同步等待的组
	wg.Add(2)

	go func() {
		for {
			res, err := server.Recv()
			if err != nil {
				wg.Done()
				break
			}
			fmt.Println(res.GetName(), res.GetAge())
		}
	}()

	go func() {
		i := 0
		for {
			if i > 10 {
				wg.Done()
				break
			}
			server.Send(&person.PersonRes{
				Name: "这是服务返回的数据",
				Age:  int32(i),
			})
			i++
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {

	// 创建监听
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
		return
	}

	// 创建服务
	s := grpc.NewServer()

	// 注册服务
	person.RegisterSearchServiceServer(s, &PersonService{})

	// 启动服务
	if err := s.Serve(l); err != nil {
		panic(err)
		return
	}
}