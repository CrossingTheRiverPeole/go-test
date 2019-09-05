package __chapterseven_interface

import (
	"fmt"
	"testing"
)

//类型分支---批量判断空接口中变量中的类型

//使用分支判断基本类型

func PrintType(v interface{})  {
	switch v.(type) {
	case int:
		fmt.Println("int 类型")
	case string:
		fmt.Println("string 类型")
	case bool:
		fmt.Println("bool 类型")
	}
}

func TestBasicType(t *testing.T)  {
	PrintType(1)
}


//----------使用类型分支判断接口类型------------------------
type CantainCanUseFaceID interface {
	canUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

type Alipay struct {

}

type Cash struct {

}

func (a *Alipay) canUseFaceID()  {

}

func (c *Cash) Stolen()  {
}

//打印支付方式具备的特点
func print(payMethod interface{})  {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceid", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen", payMethod)
	}
}

func TestInterfaceType(t *testing.T)  {
	p := new(Alipay)
	print(p)
	print(new(Cash))
}

