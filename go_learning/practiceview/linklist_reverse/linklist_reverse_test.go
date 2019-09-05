package linklist_reverse

import (
	"fmt"
	"testing"
)

type Person struct {
	Id   int
	Name string
}

type Node struct {
	Val  Person
	Next *Node
}

//遍历链表
func IteratorLinkListNode(l *Node)  {
	for   {
		fmt.Println(l.Val)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
}

//链表翻转
func LinkListReverse(p *Node) *Node {
	var  prev *Node = nil
	for p != nil {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return prev
}

func TestLinkListReverse(t *testing.T) {
	p1 := Person{
		Id:   1,
		Name: "a",
	}
	p2 := Person{
		Id:   2,
		Name: "b",
	}
	l := Node{
		Val: p1,
		Next: &Node{
			Val:  p2,
			Next: nil,
		},
	}
	IteratorLinkListNode(LinkListReverse(&l))
}
