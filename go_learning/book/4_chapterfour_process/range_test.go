package __chapterfour_process

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T)  {
	str := "hello , 你好"
	for index,value := range str {
		fmt.Println(index,value)
	}

}
