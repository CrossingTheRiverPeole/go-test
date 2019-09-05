package test

import (
	"fmt"
	"testing"
)

type Peoples interface {
	Speak(string) string
}

type Students struct {
}

func (stu *Students) Speak(think string) (talk string)  {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func live() Peoples{
	var stu *Students
	return stu
}

func TestSpeak(t *testing.T)  {
	var peo = new(Students)
	think := "bitch"
	fmt.Println(peo.Speak(think))

	var stu interface{}
	if stu == nil {
		fmt.Println("student is nil")
	}


	if live() == nil{
		fmt.Println("aaaa")
	}else{
		fmt.Println(live())
		fmt.Println("bbbbb")
	}

}


