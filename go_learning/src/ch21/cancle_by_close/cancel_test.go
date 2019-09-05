package cancle_by_close

import (
	"fmt"
	"testing"
	"time"
)

/**
通过channel来判断是否取消任务：当channel关闭的时候，就无法收到通道发出的数据，通过此机制来判断是否来关闭channel
 */

 //关闭通道
func Cancel(ch chan struct{})  {
	close(ch)
}

//判断
func isCancel(ch chan struct{}) bool  {
	select {
	case str := <- ch:
		fmt.Println(str)
		fmt.Println("执行 true")
		return true
	default:
		fmt.Println("执行 false")
		return false
	}
}
 
func TestCancelByCloseChan(t *testing.T)  {
	//启动五个协程，当channel关闭的时候来关闭携程（即协程执行完毕）
	ch := make(chan struct{})
	for i := 0; i < 5 ; i++  {
		go func(i int, ch chan struct{}) {
			for  {
				if  isCancel(ch){
					fmt.Println("任务正在执行")
				}else{
					break
				}
			}
			fmt.Println(i,"任务取消")
		}(i, ch)
	}
	time.Sleep(time.Second * 3)
}

func TestSelect(t *testing.T)  {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		for   {
			if isCancel(ch){
				fmt.Println("执行了")
				break
			}
		}
	}(ch)

	time.Sleep(time.Second * 1)
	fmt.Println("111")
	Cancel(ch)
	fmt.Println("222")
	time.Sleep(time.Second * 3)
}
