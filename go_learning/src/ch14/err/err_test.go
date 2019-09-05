package err

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)
/**
错误的处理：在函数中可以返回error，在go语言中是没有异常机制的，error可以自定义，使用errors.new("")创建自己的错误
另外，可以把无错误定义为常量，对错误进行分类，
程序中在处理错误的时候，有一种及早失败原则避免代码多层嵌套，
实现斐波那契数列结合错误处理机制
 */
func Fei(a int) ([]int, error) {
	if a < 2 || a > 100 {
		return nil, errors.New("a is not a good")
	}
	fibList := []int{1, 1}

	for i := 2; i < a; i ++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}
func TestErr(t *testing.T) {
	fmt.Println(Fei(5))
}
