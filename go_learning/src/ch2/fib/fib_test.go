package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	a := 1
	b := 1

	t.Log(a)
	for i := 1; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = a + tmp
	}

}

func TestSwap(t *testing.T)  {
	//交换两个值
	 a := 1
	 b := 2
	 a, b = b ,a


}
