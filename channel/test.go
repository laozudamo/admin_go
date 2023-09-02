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
	//
	//s := Student{
	//	Class: "21",
	//	User: User{
	//		Name: "tom",
	//		Age:  20,
	//	},
	//}
	//reflectFunc(&s)

	//TheSync()
	//theCond()
	testLock()

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
	//t := reflect.TypeOf(inter)
	//ty := t.Kind()
	//
	//switch ty {
	//case reflect.Struct:
	//	fmt.Println("asa")
	//case reflect.Pointer:
	//	fmt.Println("pointer")
	//default:
	//	fmt.Println("都不是")
	//}
	v := reflect.ValueOf(inter)
	e := v.Elem()
	e.FieldByName("Name").SetString("jack")
	fmt.Println(v)
}

/*
sync 锁

// 记住所有的锁都使用地址
*/
func TheSync() {
	// 互斥锁
	//l := &sync.Mutex{}
	//var wg sync.WaitGroup
	//wg.Add(2)
	// 读写锁
	//l := &sync.RWMutex{}

	//go lockfun(l, &wg)
	/*
		lock.Lock() 会阻塞，直到获取到锁
		lock.Unlock() 会释放锁
	*/
	//go lockfun(l, &wg)

	//wg.Wait()
	//
	//go ReadLockFun(l)
	//
	//go ReadLockFun(l)
	//
	//time.Sleep(time.Second * 10)

	// once只读写一次
	//o := &sync.Once{}
	//for i := 0; i < 10; i++ {
	//	o.Do(func() {
	//		fmt.Println("do once")
	//	})
	//}
	syncMap()
}

/*
使用的是同一把锁，所以会阻塞 互斥锁
*/
func lockfun(lock *sync.Mutex, wg *sync.WaitGroup) {
	lock.Lock() // 写的时候会排斥其他锁
	fmt.Println("lockfun")
	time.Sleep(time.Second)
	lock.Unlock()
	wg.Done()
}

func ReadLockFun(lock *sync.RWMutex) {
	lock.RLock() // 在读取的时候 不会阻塞其他的读锁 但是在写的时候会排斥其他锁
	fmt.Println("lockfun")
	time.Sleep(time.Second)
	lock.RUnlock()
}

// map 并发字典 store, load, delete, loadOrStore, range

func syncMap() {
	//var wg sync.WaitGroup
	//wg.Add(2)

	//m := sync.Map{}
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		m.Store(i, i)
	//	}
	//	wg.Done()
	//}()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(m.Load(i))
	//	}
	//	wg.Done()
	//}()

	m := &sync.Map{}

	//go func() {
	//delete(m, "1")
	m.Store("1", "1")
	m.LoadOrStore("3", "2")
	//wg.Done()
	//}()

	//go func() {
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(time.Second)
		return true
	})
	//wg.Done()
	//}()

	//wg.Wait()
}

// 并发池
func thePool() {
	p := &sync.Pool{}
	p.Put("1")
	p.Put("2")
	p.Put("3")
	p.Put("4")
	p.Put("5")
	for i := 0; i < 6; i++ {
		fmt.Println(p.Get())
	}
}

// cond

func theCond() {
	co := sync.NewCond(&sync.Mutex{})
	//co.L.Lock()
	//co.Wait()
	//co.Wait()
	//co.Wait()
	//co.L.Unlock()
	//co.Signal() // 通知co完成一个wait
	//co.Broadcast() // 通知co完成所有的wait

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		co.L.Lock()
		fmt.Println("lock1")
		co.Wait()
		fmt.Println("unlock1")
		co.L.Unlock()
		wg.Done()
	}()

	go func() {
		co.L.Lock()
		fmt.Println("lock2")
		co.Wait()
		fmt.Println("unlock2")
		co.L.Unlock()
		wg.Done()
	}()

	co.Broadcast()
	wg.Wait()
}

func testLock() {
	l := &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		l.Lock()
		time.Sleep(time.Second * 1)
		fmt.Println("lock1")
		l.Unlock()
		wg.Done()
	}()
	go func() {
		l.Lock()
		time.Sleep(time.Second * 1)
		fmt.Println("lock2")
		l.Unlock()
		wg.Done()
	}()
	fmt.Println("jack")

	wg.Wait()

}
