package main

import (
	"admin_go/grpc_class/new_pb/person"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//streamIn()
	//streamOut()

	streamInOut()
}

func streamIn() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := person.NewSearchServiceClient(conn)

	// 流式调用
	c, err := client.SearchIn(context.Background())
	if err != nil {
		panic(err)
	}

	i := 0
	for {
		if i > 10 {
			// 结束,接收服务端数据
			res, err := c.CloseAndRecv()
			if err != nil {
				panic(err)
			}

			fmt.Println(res.GetName(), res.GetAge())
			break
		}

		time.Sleep(time.Second)

		c.Send(&person.PersonReq{
			Name: "我是发送的数据",
			Age:  0,
		})
		i++
	}

}

func streamOut() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	client := person.NewSearchServiceClient(conn)

	c, err := client.SearchOut(context.Background(), &person.PersonReq{Name: "search out 信息", Age: 0})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		res, err := c.Recv()
		if err != nil {
			fmt.Println("没有其他数据接收到了,退出去")
			break
		}
		fmt.Println(res.GetName(), res.GetAge())
	}
}

// 流式调用
func streamInOut() {
	var wg sync.WaitGroup
	wg.Add(2)

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	client := person.NewSearchServiceClient(conn)

	c, err := client.SearchInOut(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			err := c.Send(&person.PersonReq{
				Name: "这是发送的信息",
				Age:  1,
			})
			if err != nil {
				wg.Done()
				break
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				wg.Done()
				break
			}
			fmt.Println(res.GetName(), res.GetAge())
		}
	}()

	wg.Wait()
}
