package __chaptersix_struct

import (
	"fmt"
	"testing"
)

//结构体内嵌
type BasicColor struct {
	R, G, B float32 // 红绿蓝三颜色
}

type color struct {
	//将基本颜色作为成员
	Basic BasicColor
	//透明度
	Alpha float32
}

func TestEmbeded(t *testing.T) {
	var c color
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0
	c.Alpha = 1

	//打印整个结构体的内容
	fmt.Printf("%+v", c)
}

// 使用组合对象描述对象特性
type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println(" can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Flying
	Walkable
}

func TestCombination(t *testing.T) {

	b := new(Bird)
	b.Fly()
	b.Walk()

	h := new(Human)
	h.Walk()
}

//初始化结构体内嵌
type Wheel struct {
	Size int
}

//引擎
type Engine struct {
	Power int
	Type  string
}

//车
type Car struct {
	Wheel
	Engine
}

func TestInitEmbedded(t *testing.T) {
	c := Car{
		Wheel: Wheel{
			Size: 4,
		},
		Engine: Engine{
			Power: 143,
			Type:  "1.4T",
		},
	}
	fmt.Printf("%+v", c)
}

//匿名结构体内嵌
type Car2 struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func TestAnonmousEmbeddedStruct(t *testing.T) {
	c := Car2{
		Wheel: Wheel{
			Size: 2,
		},
		Engine: struct {
			Power int
			Type  string
		}{Power: 3, Type: "1.21"},
	}
	fmt.Printf("%+v",c)
}
