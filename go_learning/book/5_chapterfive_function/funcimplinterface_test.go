package __chapterfive_function

import (
	"fmt"
	"testing"
)

//函数实现接口
// 函数实现接口时不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体
type invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func TestFuncImplInterface(t *testing.T) {
	var invoker invoker

	invoker = FuncCaller(func(p interface{}) {
		fmt.Println(p)
	})
	invoker.Call("hello")
}