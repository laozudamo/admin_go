package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func main() {
	//handleChan()
	//u := User{
	//	Name: "jack",
	//	Age:  18,
	//}
	//Check(u)
	//check(u)

	s := Student{
		Class: "21",
		User: User{
			Name: "tom",
			Age:  20,
		},
	}
	reflectFunc(&s)

}

func testBinfa() {
	var wg sync.WaitGroup
	wg.Add(1)

	go Run(&wg) // 开启一个goroutine执行任务
	for i := 0; i < 10; i++ {
		fmt.Println("main----", i)
	}
	wg.Wait() // 阻塞，直到计数器变为0
}

func Run(wg *sync.WaitGroup) {
	fmt.Println("run----")
	wg.Done() // 任务完成，减少一个计数
}

func TestChan() {
	ch1 := make(chan int, 5) // 有缓冲的channel,存储完数据后
	//ch1 := make(chan int) // 无缓冲的channel, 先到39 -> 33

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}

	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
}

func ReadChan() {
	c1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
		wg.Done()
	}()
	for v := range c1 {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
	wg.Wait()
}

func handleChan() {
	c := make(chan int)
	var wc chan<- int = c
	var rc <-chan int = c
	go SetChan(wc)
	GetChan(rc)
}

func SetChan(wc chan<- int) {
	for i := 0; i < 10; i++ {
		wc <- i
	}
	close(wc)
	fmt.Println("set 函数")
}

func GetChan(rc <-chan int) {
	for v := range rc {
		fmt.Println(v)
	}
	fmt.Println("get 函数")

}

func handleChans() {
	c := make(chan int)
	go SetChan(c)
	GetChan(c)

}

type User struct {
	Name string
	Age  int
}

type Student struct {
	Class string
	User
}

// 为User类型绑定一个方法
func (u User) sayName() {
	fmt.Println(u.Name)
}

func Check(v interface{}) {
	user, ok := v.(User)

	user.sayName()

	// 第一个参数返回对象，第二个参数返回是否成功
	fmt.Println(user, ok)
}

// 类型断言
func check(v interface{}) {
	switch v.(type) {
	case User:
		fmt.Println("user", v.(User).Name)
	case Student:
		fmt.Println("student", v.(Student).Class, v.(Student).Name)
	}
}

// 反射
func reflectFunc(inter interface{}) {
	t := reflect.TypeOf(inter)
	ty := t.Kind()

	switch ty {
	case reflect.Struct:
		fmt.Println("asa")
	case reflect.Pointer:
		fmt.Println("pointer")
	default:
		fmt.Println("都不是")
	}
}
