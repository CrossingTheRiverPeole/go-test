package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
多个任务返回结果，发送到channel，当拿到channel中的一个结果的时候，即可返回结果，注意，当使用channel接收多个goroutine返回的结果的时候，
注意为channel设置缓存空间，如果不设置缓存空间，当多个goroutine发送结果的时候，只要一个结果能放到goroutine中，后续的会阻塞，这样goroutine会逐渐增多
 */


func RunTask(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("The result is from %d", i)
}

func FirstResponse() string {
	numberOfRunner := 10
	// 注意当只需从chan中获取一个结果的时候，注意设置一个缓存空间，如果不设置缓存空间，剩下的goroutine会阻塞，因为没法向channel中写入数据
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i++ {
		go func(i int) {
			ch <- RunTask(i)
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
