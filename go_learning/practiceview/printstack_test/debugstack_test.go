package printstack_test

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestDebugStack(t *testing.T)  {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			fmt.Println("recover")
		}
	}()
	panic("new error")
}
