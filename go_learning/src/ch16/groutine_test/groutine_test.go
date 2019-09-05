package groutine_test

import (
	"fmt"
	"testing"
)
func TestGroutine(t *testing.T)  {
	//注意：使用goroutine时注意传值时是下面的写法，而不是直接调用循环
	for i := 0; i< 10; i ++{
		go func(i int) {
			fmt.Println("i的值", i )
		}(i)
	}
}
