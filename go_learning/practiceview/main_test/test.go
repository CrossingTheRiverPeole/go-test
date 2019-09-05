package main

import (
	"fmt"
)

func modifySlice(s []string)  {
	s[0] = "111"
	fmt.Printf("%p",&s)
	s = append(s, "456")
	fmt.Println("")
	fmt.Printf("%p",&s)
}


func TestSlice()  {

}


type Student struct {
	name string
	age int
}



func main()  {
	/*s := make([]string, 4,4)
	s[0] = "123"
	modifySlice(s)
	fmt.Println(s)*/

	s := []Student{
		{name: "tom", age:11},
		{name:"jack", age:12},
		{name:"song", age:13},
	}

	m := make(map[string]*Student)
	//遍历
	for _,v := range s{
		m[v.name] = &v
		fmt.Println("")
		fmt.Printf("%p", &v)
		fmt.Println(&v)
	}
	for i,v := range s{
		m[v.name] = &s[i]
		fmt.Println("")
		fmt.Printf("%p", &s[i])
	}

	for k, v := range m{
		fmt.Println(k)
		fmt.Println(*v)
	}
}
