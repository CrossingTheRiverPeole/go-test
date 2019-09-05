package __chapterthree_container

import (
	"container/list"
	"fmt"
	"testing"
)

//list是一个非连续的存储空间,变量记录彼此之间的联系
func TestList(t *testing.T) {
	//两种初始化方法：一种是通过list.new()一种是通过声明
	l := list.New()
	var l2 list.List

	l.PushFront(1)
	l2.PushBack(1)

	fmt.Println(l)
	fmt.Println(l2)
}



//list方法操作
func TestListOperate(t *testing.T) {
	l := list.New()

	l.PushFront("first")
	el := l.PushBack("last")
	l.InsertBefore(2, el)
	l.InsertAfter(3, el)

	//遍历
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
