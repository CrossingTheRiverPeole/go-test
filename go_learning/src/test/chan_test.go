package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestChan(t *testing.T)  {

	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-string_chan:
		panic(value)
	case value := <-int_chan:
		fmt.Println(value)

	}
}

func TestSliceAppend(t *testing.T)  {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1,s2...)
	fmt.Println(s1)



}




