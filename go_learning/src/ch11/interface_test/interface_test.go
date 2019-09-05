package interface_test

import (
	"fmt"
	"testing"
)

type Inner struct {
}

/**
接口的实现与接口是不相互依赖的，实现接口之后，
把实现接口的结构体重新赋值给接口，然后调用接口
 */
// 定义接口
type Lost interface {
	add() bool
}

//实现接口
func (inner *Inner) add() bool {
	fmt.Println("添加失败")
	return false
}

func TestInterface(t *testing.T) {
	inner := new(Inner)
	//lost := inner
	//inner.add()
	var lost Lost
	lost = inner
	lost.add()
}
