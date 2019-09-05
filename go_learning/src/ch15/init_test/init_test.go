package init_test

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("2222init方法2")
}
func init() {
	fmt.Println("11111init方法执行")
}

/**
init方法可以在一个文件中定义多个，而且是按照定义的顺序执行
 */
func TestInit(t *testing.T) {
	fmt.Println("init方法测试")
}
