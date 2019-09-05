package __chapterfour_process

import (
	"fmt"
	"testing"
)

//跳出多层循环
func TestBreak(t *testing.T) {
	//OutLoop
OutLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(j)
			switch j {
			case 7:
				fmt.Println(j)
				break OutLoop
			}

		}
	}
}

//继续下一循环:使用continue跳出本次循环执行下一次训话
func TestContinue(t *testing.T) {
OutLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(j)
			switch j {
			case 2:
				fmt.Println("continue next loop")
				continue OutLoop
			}
		}
	}
}

//使用goto跳出循环
func TestGoto(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakHere
			}
		}
	}

breakHere:
	fmt.Println("跳出循环")
}
