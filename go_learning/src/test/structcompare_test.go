package test

import (
	"fmt"
	"testing"
)

func TestStructTest(t *testing.T)  {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	//sm1与sm2是不能比较的，结构体中有引用类型是不能比较的，如channel、map、slice
	fmt.Println(sm1,sm2)

	i := 1
	i++
	fmt.Println(i)



}
