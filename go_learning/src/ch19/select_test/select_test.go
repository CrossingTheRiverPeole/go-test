package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "done"
}

func AsyncService() chan string {
	retch := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result")
		retch <- ret
		fmt.Println("service exited")
	}()
	return retch
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 1000):
		fmt.Println("时间超时")
	}
}
