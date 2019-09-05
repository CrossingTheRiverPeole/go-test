package __chapterfive_function

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println("defer begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("defer end")
}

//使用延迟并发解锁
var (
	valueKey        = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

func TestDeferMetux(t *testing.T) {
	key := "test"
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()
	v := valueKey[key]
	fmt.Println(v)
}

//使用延迟释放文件句柄
func fileSize(filename string) int64 {
	//根据文件名打开文件
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return 0
	}

	//获取文件状态信息
	info, err := f.Stat()
	//获取文件状态信息时发生错误，关闭文件并返回文件大小为0
	if err != nil {
		return 0
	}

	//获取文件大小
	size := info.Size()
	return size
}










