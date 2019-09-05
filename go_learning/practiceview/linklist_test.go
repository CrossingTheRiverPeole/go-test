package practiceview

import (
	"fmt"
	"testing"
)

type Student struct {
	Id   int
	Name string
}

type Node struct {
	Val  Student
	Next *Node
}

//遍历链表
func (l *Node) IteratorLinkList() {
	if l == nil {
		return
	}
	for {
		fmt.Println(*l)
		if l.Next == nil {
			break
		}
		l = l.Next
	}

}

func TestLinkList(t *testing.T) {
	v1 := Student{
		Id:   1,
		Name: "A",
	}
	v2 := Student{
		Id:   2,
		Name: "B",
	}

	linkList := &Node{
		Val: v1,
		Next: &Node{
			Val: v2,
		},
	}
	linkList.IteratorLinkList()
}
