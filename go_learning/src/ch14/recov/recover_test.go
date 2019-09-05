package recov

import (
	"errors"
	"fmt"
	"testing"
)

/**
panic vs os.Exit
os.Exit退出时不会调用defer指定的函数
os.Exit退出时不输出当前调用栈的信息
 */
func TestRecover(t *testing.T) {
	//使用recover恢复
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()
	//defer是在panic之前执行的
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	panic(errors.New("panic create"))
}
