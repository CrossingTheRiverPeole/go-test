package __chapter_concurrent

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
)

//原子访问包（atomic）、互斥锁sync.Metuex以及等待锁（waitGroup）


//序列号
var seq int64
//原子访问包
func GenID() int64  {
	return atomic.AddInt64(&seq, 1)
}
func TestAtomic(t *testing.T)  {
	for i := 0; i < 10 ; i++  {
		go GenID()
	}

	fmt.Println(GenID())

}

//-----------------互斥锁：sync.Mutex--------------

var(
	// 定义变量
	count int
	// 定义锁
	lock sync.Mutex
)

func GetCount()int  {
	lock.Lock()
	defer  lock.Unlock()
	return count
}

func SetCount(c int)  {
	lock.Lock()
	defer lock.Unlock()
	count  = c
}

func TestSyncMutex(t *testing.T)  {
	SetCount(9)
	fmt.Println(GetCount())
}

//--------------读写锁：用在读比写多的情况，比互斥锁更高效---------------
var(
	c int
	countGuard sync.RWMutex
)

func GetCountByRWMutex() int {
	countGuard.Lock()
	defer countGuard.Unlock()

	return c
}


//-------------等待组sync.waitGroup------------------------
func TestWaitGroup(t *testing.T)  {
	var wg sync.WaitGroup

	urls := []string{
		"http://www.baidu.com",
		"https://www.github.com",
		"https://www.golangtc.com",
	}

	for _, url := range urls {
		wg.Add(1)
		//开启一个并发
		go func(url string) {
			defer wg.Done()
			//访问网址
			_,err := http.Get(url)
			fmt.Println(url,err)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}




