package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type Student struct {
	Name string
	Age int
}

func TestCode(t *testing.T)  {

	m := make(map[string]*Student)

	stus := []Student{
		{Name:"song",Age:12},
		{Name:"lig", Age:13},
	}

	for _, stu := range stus  {
		fmt.Println(stu)
		m[stu.Name] = &stu
	}

	//for key,v := range m  {
	//	fmt.Println(key, v)
	//}


}

func TestPrintInt(t *testing.T)  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i=: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestDefer(t *testing.T)  {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
