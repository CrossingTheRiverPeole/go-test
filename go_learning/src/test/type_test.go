package test

import (
	"fmt"
	"testing"
)

func GetValue() interface{} {
	return 1
}

func TestType(t *testing.T) {
	i := GetValue()
	// type只能用在interface上
	switch i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	default:
		fmt.Println("unknown type")
	}
	
}
