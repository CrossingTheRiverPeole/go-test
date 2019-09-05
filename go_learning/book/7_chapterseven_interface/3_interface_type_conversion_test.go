package __chapterseven_interface

import (
	"fmt"
	"testing"
)

//使用接口断言将接口转换成另一个接口，也可以将接口转换为另外的类型。
type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type Bird struct {

}

type Pig struct {

}

func (b *Bird) Fly(){
	fmt.Println("bird： fly")
}

func (b *Bird) Walk()  {
	fmt.Println("bird: walk")
}

func (p *Pig) Walk()  {
	fmt.Println("pig: walk")
}

func TestInterfaceAssert(t *testing.T)  {
	m := map[string]interface{}{
		"bird": new(Bird),
		"pig": new(Pig),
	}

	for _, v :=range m  {
		bird, isFlyer := v.(Flyer)
		pig, isPig := v.(Walker)
		if isFlyer{
			bird.Fly()
		}
		if isPig {
			pig.Walk()
		}
	}
}




