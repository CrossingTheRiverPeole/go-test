package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	x := 1
	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}

func TestRoutinePrint(t *testing.T) {
	str := []string{"one", "two", "three"}

	for _, s := range str {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println(s)
		}()
	}
	time.Sleep(time.Second * 2)
}

func TestSliceIndex(t *testing.T) {
	str := []string{"aa", "bb", "cc"}
	for v := range str {
		fmt.Println(v)
	}

}

type Slice []int

func (s *Slice) add(a int) *Slice {
	*s = append(*s, a)
	fmt.Println(a)
	return s
}

func TestDefer2(t *testing.T) {

	m := make(Slice, 1)
	defer m.add(1).add(2)

	m.add(3)

}

func TestFileOpen(t *testing.T) {
	file, err := os.Open("test.go")

	defer file.Close()

	if err != nil {

		fmt.Println("open file failed:", err)

		return
	}
}

func TestChar(t *testing.T)  {
	str := "hello"
	fmt.Println(str)

}
