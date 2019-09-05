package __chaptersix_struct

import (
	"fmt"
	"testing"
)

//定义结构体
type Point struct {
	X int
	Y int
}

//实例化结构体
func TestStructDefine(t *testing.T)  {
	var p Point
	p.X = 2
	p.Y = 3
	fmt.Printf("%+v", p)
}

type Player struct {
	Name string
	HealthPoint int
	MagicPoint int
}
// 创建指针类型的结构体
func TestPointerStruct(t *testing.T)  {
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
}

type Command struct {
	Name string
	Var *int
	Comment string
}

func TestStructInstantiation(t *testing.T)  {
	cmd := &Command{}
	cmd.Name = "version"
}

// 初始化结构体成员
type People struct {
	name string
	child *People
}

func TestInitStruct(t *testing.T)  {

	relation := &People{
		name:"爷爷",
		child: &People{
			name:"爸爸",
			child: &People{
				name:"我",
			},
		},
	}
	fmt.Printf("%+v", relation)
}

//初始化匿名结构体
func printMsg(msg *struct{
	id int
	data string
})  {
	fmt.Printf("%T", msg)
}

func TestAnonmousStruct(t *testing.T)  {
	msg := &struct {
		id int
		data string
	}{
		id: 1024,
		data: "hello",
	}
	printMsg(msg)
}

//构造函数--结构体和类型的一系列初始化操作的函数封装

type Cat struct {
	Name string
	Color string
}

func NewCatByName(name string) *Cat  {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string)*Cat  {
	return &Cat{
		Color: color,
	}
}

	//带有父子关系的结构体的构造和初始化--模拟父级构造调用
	type BlackCat struct {
		cat Cat
	}
	//构造基类
	func NewCat(name string) *Cat  {
		return &Cat{
			Name: name,
		}
	}
	//构造子类
	func NewBlackCat(color string) *BlackCat {
		cat := &BlackCat{}
		cat.cat.Color = color
		return cat
}







