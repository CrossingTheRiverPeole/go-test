package main

import "fmt"

func Escape() []func() {
	f := make([]func(),3)
	for i := 0; i< 3 ; i++  {
		f[i] = func() {
			fmt.Print(i)
		}
	}
	return f
}

func main()  {
	f := Escape()

	f[0]()
	f[1]()
	
}

