package implments

import (
	"fmt"
	"testing"
)
type code string
type Programmer interface {
	WriteHelloWorld() code
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() code {
	return "java hello world"
}

type GoProgramer struct {
}

func (g GoProgramer) WriteHelloWorld() code {
	return "go hello world"
}

func TestImplement(t *testing.T) {
	j := new(JavaProgrammer)
	g := new(GoProgramer)

	fmt.Println(j.WriteHelloWorld())
	fmt.Println(g.WriteHelloWorld())




}
