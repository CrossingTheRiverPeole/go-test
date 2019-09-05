package _0_chapter_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

//reflect.Type获取值的类型信息
func TestReflectType(t *testing.T)  {
	var a int
	typeOf := reflect.TypeOf(a)
	fmt.Println(typeOf.Name(),typeOf.Kind())
}

func TestReflectStruct(t *testing.T)  {

	type Cat struct {
		Name string
		Type int `json:type id:"100"`
	}
	//创建cat实例
	ins := Cat{
		Name: "mimi",
		Type: 1,
	}

	typeOfCat := reflect.TypeOf(ins)
	//遍历所有的结构体成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		fildType := typeOfCat.Field(i)
		fmt.Printf("name: %v tag: '%v'\n", fildType.Name, fildType.Tag)
	}

	if catType, ok := typeOfCat.FieldByName("Type"); ok{
		//从tal中取出需要的信息
		fmt.Println(catType.Tag.Get("json"),catType.Tag.Get("id"))
	}
}
//----------------tag------------------
func TestPointer(t *testing.T)  {
}

//通过反射修改值
func TestReflectValueof(t *testing.T)  {
	a := 1024
	valueOfA := reflect.ValueOf(&a)
	valueOfA = valueOfA.Elem()
	valueOfA.SetInt(3)

	fmt.Println(valueOfA.Int())
}



