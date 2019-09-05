package definetype_test

import (
	"fmt"
	"testing"
	"time"
)

/**
测试自定义类型
 */

//自定义类型
type opfn func(int) int
func total(op opfn) opfn {
	return func(i int) int {
		start := time.Now().Second()
		fmt.Println(op(i))
		return time.Now().Second() - start
	}
}
func sum(op int) int {
	var sum int
	for i := 0; i < 10; i++ {
		sum += op
	}
	return sum
}
func TestDefineType(t *testing.T) {
	t.Log(total(sum)(5))
}
