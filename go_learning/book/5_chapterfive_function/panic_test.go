package __chapterfive_function

import (
	"fmt"
	"testing"
)

//panic与defer
func TestPanic(t *testing.T)  {
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	defer func() {
		recover()
	}()
	panic("宕机")
}


