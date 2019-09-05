package struct_test

import (
	"testing"
)

type Employee struct {
	Id int
	Name string
}

func TestCreatingStr(t *testing.T)  {
	e := Employee{
		Id: 1,
		Name: "song",
	}

	e2 := new(Employee)
	e2.Id = 2
	e2.Name = "song"

	t.Logf("%T", e)
	t.Logf("%T", e2)
}
