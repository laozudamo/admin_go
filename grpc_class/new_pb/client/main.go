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

func streamInOut() {

	var wg sync.WaitGroup // 定义一个同步等待的组
	wg.Add(2)             // 2个任务

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
		i := 1
		for {
			if i > 10 {
				break
				defer wg.Done()
			}

			c.Send(&person.PersonReq{
				Name: "我是传递过来的数据",
				Age:  int32(i),
			})
			i++
			time.Sleep(time.Second)
		}

		fmt.Println("协程1执行")

	}()

	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				break
				defer wg.Done()
			}
			fmt.Println(res.GetName(), res.GetAge())
		}

		fmt.Println("协程2执行")
	}()

	wg.Wait() // 等待所有协程执行完毕
}
