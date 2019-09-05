package share_mem_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestShareMem(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	t.Log(counter)
}

/**
使用锁保证counter的计数
 */
func TestShareMemThreadSafe(t *testing.T) {
	var met sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				met.Unlock()
			}()
			met.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

/**
使用waitgroup保证协程的执行顺序
 */
func TestShareMemWaitGroup(t *testing.T)  {
	var met sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i :=0; i< 5000 ; i ++ {
		wg.Add(1)
		go func() {
			defer func() {
				met.Unlock()
			}()
			met.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
}

func TestWatiGroup(t *testing.T)  {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 1 ; i++  {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

