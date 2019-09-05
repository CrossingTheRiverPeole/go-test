package __chaptersix_struct

import (
	"fmt"
	"testing"
)

//指针类型接收器

// 定义结构体
type Property struct {
	value int
}

func (p *Property) SetValue( v int)  {
	p.value = v
}

func (p *Property) GetValue() int {
	return p.value
}

func TestPointReciver(t *testing.T)  {
	p := &Property{}
	p.SetValue(8)
	fmt.Println(p.GetValue())
}

//非指针接收器方法
type Pointer struct {
	X int
	Y int
}

//非指针接收器
func (p Pointer) Add(other Pointer) Pointer  {
	return Pointer{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func TestUnpointerReciver(t *testing.T)  {
	p1 := Pointer{1,1}
	p2 := Pointer{2,2,}
	fmt.Println(p1.Add(p2))
}


