package all_response_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
获取所有goroutine的执行结果：等待协程把结果都写入到channel中，如果当前协程未完全写入执行结果，那从channel读取的时候会阻塞，直到所有结果放入
与只读取一个结果相比，只读取一个结果的时候，只要拿到一个结果就返回，这样就OK了，不过这样的话也要注意给channel一个缓存空间，这样其余剩下的goroutine
执行完之后才能把结果放入而不是一直阻塞在那里
 */


func RunTask(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("the result is from %d", i)
}

func AllResponse() string {
	numberRunner := 10
	ch := make(chan string, numberRunner)
	for i := 0; i < numberRunner; i++ {
		go func(i int) {
			ch <- RunTask(i)
		}(i)
	}

	finalRet := ""
	for i := 0; i < numberRunner; i++ {
		finalRet += <-ch
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
