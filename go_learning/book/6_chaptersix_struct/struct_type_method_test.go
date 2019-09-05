package __chaptersix_struct

import (
	"fmt"
	"testing"
)

//为类型添加方法，重新定义类型： type 【类型名】 类型

// 将int定义为MyInt类型
type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return int(m) + other
}

func TestStructTypeMethod(t *testing.T)  {
	var b MyInt
	fmt.Println(b.IsZero())
	b = 1
	fmt.Println(b.Add(1))
}




