package sync_pool_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
sync.pool是一个缓存池，里面存放需要缓存的对象，但是池中的对象受runtime.gc的影响，池中对象的生命周期是到gc回收之前，另外，这个池
是一个同步的，存在一定的性能问题，对于创建对象还是直接从池中取对象，两者要进行比较
 */
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new Object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC会清除缓存中的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

//测试多协程
func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new Object")
			return 100
		},
	}

	pool.Put(400)
	pool.Put(400)
	pool.Put(400)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(pool.Get())
			wg.Done()
		}()
	}
	wg.Wait()
}
