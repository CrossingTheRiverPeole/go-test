package test

import (
	"fmt"
)

/**
交替循环打印12AB34CD
 */
func estPrint() {

	//解题思路：两个循环分别打印数组和字母，两个channel交替向通道中写入数据，第三个通道阻塞，直到打印完成
	chanD := make(chan bool,1)
	chanC := make(chan bool)
	done := make(chan bool)

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	charSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	go func() {
		for i := 0; i < 8; i = i + 2 {
			<-chanC
			fmt.Println(intSlice[i])
			fmt.Println(intSlice[i+1])
			chanD <- true
		}
	}()

	go func() {
		for i := 0; i < 8; i = i + 2 {
			<-chanD
			fmt.Println(charSlice[i])
			fmt.Println(charSlice[i+1])
			chanC <- true
		}
		done <- true
	}()
	chanC <- true
	<-done
}
