package __chapterfour_process

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {

	//几种不同的用法：
	//第一种：基本写法
	str := "hello"
	switch str {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("over")
	}
	//一分支多值
	a := "mum"
	switch  a{
	case "mum", "daddy":
		fmt.Println("mum daddy")
	default:
		fmt.Println("default")
	}




}
