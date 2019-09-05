package fn_test

import (
	"fmt"
	"testing"
	"time"
)

// 函数式编程：函数作为参数传递进去，并返回一个函数
func funSpend(in func(int) int) func(int) int {
	return func(op int) int {
		start := time.Now()
		op = in(op)
		println("执行时间:", time.Since(start).Seconds())
		return op
	}
}

func testFn(in int) int {
	time.Sleep(time.Second * 3)
	return in
}

func TestTimeSpend(t *testing.T) {
	// 传递一个函数进去
	fn := funSpend(testFn)
	// 执行返回的函数
	fn(3)
}

/**
可变参数：参数是多个参数，默认转换成数组，然后使用range进行遍历
 */
func Sum(ops ...int) {
	s := 0
	for op := range ops {
		s += op
	}
	fmt.Println(s)
}

func TestVarParam(t *testing.T) {

	Sum(1, 2, 3, 4, 5)
}

//测试defer的执行顺序，当有多个defer的时候，确认defer的执行顺序（从后向前执行），并且在panic执行执行
func TestDefer(t *testing.T) {
	t.Log("开始执行")
	defer func() {
		t.Log("clear resources")
	}()
	panic("panic  error")
}
