package test

import (
	"fmt"
	"testing"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	p People
}

func (t *Teacher) ShowA()  {
	fmt.Println(" teacher showA")
	t.ShowB()
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestP(t *testing.T)  {
	tt := &Teacher{}
	tt.ShowA()
}
