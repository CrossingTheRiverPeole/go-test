package slice_test

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T) {
	var s0 []int
	fmt.Println(len(s0), cap(s0))
	s0 = append(s0,1)
	fmt.Println(len(s0),cap(s0))
}

func TestSlice(t *testing.T)  {
	var intSli = []int{1,2,3,4,5,6,7,8,9,10,11,12}
	sli2 := intSli[0:3]
	fmt.Println(len(sli2),cap(sli2))
}


