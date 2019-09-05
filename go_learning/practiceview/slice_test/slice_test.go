package slice_test

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
	Age int
}

func modifySlice(s []string)  {
	s[0] = "111"
	s = append(s, "456")
	fmt.Println(s)
}

func TestSlice(t *testing.T)  {
	s := make([]string, 4,4)
	s[0] = "123"
	modifySlice(s)
	fmt.Println(s)
}
