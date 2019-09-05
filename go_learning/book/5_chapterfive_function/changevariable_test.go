package __chapterfive_function

import (
	"bytes"
	"fmt"
	"testing"
)

//定义可变参数
func joinString(slist ...string) string  {
	//遍历slist，并拼接返回
	var b bytes.Buffer

	for _, str := range slist  {
		b.WriteString(str)
	}
	return b.String()
}

func TestChangeVariables(t *testing.T)  {
	fmt.Println(joinString("pig", "and", "rat"))
	fmt.Println(joinString("hammer", "mom", "and", "hawk"))
}

//获得可变参数的类型
func printTypeValye(slist ...interface{}) string  {
	// 字符串缓冲作为字符串连接
	var b bytes.Buffer
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		//类型的字符串描述
		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value:")
		b.WriteString(str)
		b.WriteString(" type")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func TestVariableType(t *testing.T)  {
	fmt.Println(printTypeValye(100, "str", true))
}



