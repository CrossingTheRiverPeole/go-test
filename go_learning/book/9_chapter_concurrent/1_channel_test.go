package __chapter_concurrent

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestChannelBlock(t *testing.T) {
	ch := make(chan int)
	go func() {
		fmt.Println("start goroutine")
		ch <- 1
		fmt.Println("exit goroutine")
	}()
	<-ch
	fmt.Println("all done")
}

func TestLoopGetData(t *testing.T) {
	ch := make(chan int)
	go func() {
		fmt.Println("start goroutine")
		for i := 0; i < 10; i++ {
			ch <- i
		}
		fmt.Println("exit goroutine")
	}()
	for data := range ch {
		fmt.Println(data)
		if data == 9 {
			break
		}
	}
	fmt.Println("all done")

}

//模拟并发打印
func Printer(c chan int) {
	for {
		data := <-c
		fmt.Println(data)
		if data == 0 {
			break
		}
	}

	c <- 0

}

func TestPrinter(t *testing.T) {

	c := make(chan int)

	go Printer(c)

	for i := 1; i < 10; i++ {
		c <- i
	}
	//通知打印，没有数据了
	c <- 0
	//等待printer结束
	<-c
}

//------------缓冲通道--------------
func TestCacheChan(t *testing.T) {
	ch := make(chan int, 2)
	//查看当前通道的大小
	fmt.Println(len(ch))

	//发送整型元素到通道
	ch <- 2
	ch <- 1

	//查看当前通道的大小
	fmt.Println(len(ch))
}

//通道多路复用select
//格式：使用多路复用模拟远程过程调用
func RPCClient(ch chan string, req string) (string, error) {
	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("time out")
	}
}

func RPCServer(ch chan string) {
	for {
		//接收客户端请求
		data := <-ch
		fmt.Println("server recived:", data)
		//向客户端反馈已经收到
		ch <- "roger"
	}
}
func TestSendRecive(t *testing.T) {
	//创建一个无缓冲的通道
	ch := make(chan string)
	//并发执行服务器逻辑
	go RPCServer(ch)
	//客户端请求数据和接收数据
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		//发生错误进行打印
		fmt.Println(err)
	} else {
		fmt.Println("client recive", recv)
	}
}

//--------------一段时间之后time.After()------------
func TestTimeAfterFunc(t *testing.T)  {
	ch := make(chan int)
	//
	fmt.Println("start")

	//1之后调用匿名函数
	time.AfterFunc(time.Second, func() {
		//一秒之后打印结果
		fmt.Println("one second after" )
		ch <- 1
	})
	<- ch
}

func TestTimeTickerAndNewTimer(t *testing.T)  {
	//创建一个打点计时器，每500毫秒执行一次
	ticker := time.NewTicker(time.Millisecond * 500)
	//创建一个计时器，2秒之后执行
	timer := time.NewTimer(time.Second * 2)
	var i int
	for{
		select {
		case <- timer.C:
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}
	// 退出标签
	StopHere:
		fmt.Println("done")
}

//关闭通道：向已关闭的通道中发送数据会触发panic，从已关闭的通道中获取数据没有问题
func TestCloseChan(t *testing.T)  {
	//创建一个缓冲通道，关闭缓冲通道，从关闭的缓冲通道中获取数据
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	//遍历通道
	for i := 0; i < cap(ch); i++  {
		v, ok := <- ch
		fmt.Println(v,ok)
	}
}

func add(a,b int) int {
	return a+b
}
//使用反射调用函数
func TestReflectFunc(t *testing.T)  {
	funValue := reflect.ValueOf(add)
	paramList := []reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}
	retList := funValue.Call(paramList)
	fmt.Println(retList[0])
}


