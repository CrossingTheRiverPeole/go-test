package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name       string `fromat:"normal"`
	Age        int
}

func (e *Employee) updataAge(age int) {
	e.Age = age
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Println("Unknown", t)
	}

}

func TestDeepEqual(t *testing.T)  {
	a := map[int]string{1: "one", 2 :"two", 3:"three"}
	b := map[int]string{1: "one", 2 :"two",  3:"three"}
	fmt.Println(reflect.DeepEqual(a,b))
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	fmt.Println(reflect.ValueOf(*e).FieldByName("Age"))

	field, _ := reflect.TypeOf(*e).FieldByName("Age")
	fmt.Println(field.Tag)
}

type User struct {
	Id int
	Name string
	Age int
}

func (u User) ReflectCallFunc()  {
	fmt.Println("call method")
}


func TestReflectUser(t *testing.T)  {
	u := User{Id:1,Name:"song", Age:12}
	//通过接口获取任意参数
	DoFileAndMethod(u)
}

func DoFileAndMethod(u interface{})  {
	getType := reflect.TypeOf(u)
	fmt.Println("get type is ", getType.Name())

	getValue := reflect.ValueOf(u)
	fmt.Println("get all fields is ", getValue)

	for i :=0 ; i < getType.NumField(); i ++  {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//获取方法
	for i := 0; i < getType.NumMethod() ;i++  {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
























