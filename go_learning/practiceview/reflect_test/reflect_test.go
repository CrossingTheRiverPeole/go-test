package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)

}

type user struct {
	Id   int
	Name string
	Age  int
}

//有参数的方法
func (u user) ReflectCallFunc(a string) {
	fmt.Println(a)
	fmt.Println("Allen .Wu RegflectCallFunc")
}

//无参数的方法
func (u user) ReflectMethodNoArgs() {
	fmt.Println("无参数")
}

func TestReflect2(t *testing.T) {

	user := user{
		Id:   1,
		Name: "song",
		Age:  2,
	}

	getType := reflect.TypeOf(user)
	getValue := reflect.ValueOf(user)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)              //获取对应的属性
		value := getValue.Field(i).Interface() //获取属性对应的值
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}
}

//获取方法
func TestReflectMethod(t *testing.T) {
	user := user{
		Id:   1,
		Name: "song",
		Age:  12,
	}

	getValue := reflect.ValueOf(user)
	//调用有参数的构造方法
	methodValue := getValue.MethodByName("ReflectCallFunc")
	args := []reflect.Value{reflect.ValueOf("song")}
	methodValue.Call(args)

	//调用无参数的构造方法
	noArgsMethod := getValue.MethodByName("ReflectMethodNoArgs")
	//无参数方法设置空参数
	noArgs := make([]reflect.Value,0)
	noArgsMethod.Call(noArgs)
}
