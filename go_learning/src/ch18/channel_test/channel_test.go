package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func service()string  {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func AsyncService() chan string  {
	retch := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result")
		retch <- ret
		fmt.Println("service exited")
	}()
	return retch
}

func otherTask()  {
	fmt.Println("workding on other else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func TestAsyncService(t *testing.T)  {
	retch := AsyncService()
	otherTask()
	fmt.Println(<- retch)
}
