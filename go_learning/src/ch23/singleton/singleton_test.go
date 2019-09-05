package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

/*
使用once.do方法实现只执行一次，保证对象只进行一次初始化
 */

type Singleton struct {
}

var singletonInstace *Singleton
var once sync.Once

func getSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("create obj")
		singletonInstace = new(Singleton)
	})
	return singletonInstace
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go func() {
			obj := getSingletonObj()
			fmt.Printf("%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

/**
使用sync.once实现单例模式
 */
var singletonObj *Singleton
var onceDone sync.Once

func CreateSingletonObj() *Singleton {
	onceDone.Do(func() {
		fmt.Println("create a object")
		singletonObj = new(Singleton)
	})
	return singletonObj
}

func TestSingletonCreate(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go func() {
			singletonObj = CreateSingletonObj()
			wg.Done()
		}()
	}
	wg.Wait()
}
