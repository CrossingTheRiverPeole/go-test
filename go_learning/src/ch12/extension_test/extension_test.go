package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}
func (p * Pet) Speak()  {
	fmt.Println("....")
}
func (p *Pet)SpeakTo()  {
	p.Speak()
	fmt.Println("wang")
}
type Dog struct {
	 *Pet
}
func (d *Dog) Speak()  {
	fmt.Println("chao")
}
func (d *Dog) SpeakTo()  {
	d.Speak()
	fmt.Println("wang方法风格分割")
}
func TestExtension(t *testing.T)  {
	d := new(Dog)
	d.SpeakTo()
}
