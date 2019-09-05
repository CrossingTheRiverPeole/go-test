package emptyinterfaceassertion

import (
	"fmt"
	"testing"
)

func  DoSomething(p interface{})  {
	switch p.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Println("type unknown")
	}
}

func TestAssert(t *testing.T)  {
	DoSomething(10)
	DoSomething("10")

}
