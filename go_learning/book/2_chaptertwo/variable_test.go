package __chaptertwo

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestType(t *testing.T) {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}

func TestString(t *testing.T) {
	var str = `第一行
   第二行
第三行
     \r\n
   `
	fmt.Println(str)

}

func TestSlice(t *testing.T) {

}

/**
取指针的值
 */
func TestPointer(t *testing.T) {
	var mode = flag.String("mode", "", "process mode")

	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*mode)

}

func TestStringLen(t *testing.T) {
	var s = "忍者"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

func TestStringRange(t *testing.T) {
	theme := "狙击 test"

	for _, s := range theme {
		fmt.Printf("%c", s)
	}
}

//取字符串中的某一段字符
func TestSubString(t *testing.T) {
 	tracer := "死神来了, 死神 bye bye"
 	comma := strings.Index(tracer, ", ")
 	pos := strings.Index(tracer[comma:], "死神")
 	fmt.Println(comma, pos, tracer[comma + pos:])
}

//修改字符串
func TestChangeString(t *testing.T) {
	angle := "heros never die"
	// 转换成byte数组
	angleByte := []byte(angle)

	for i := 5; i <= 10; i++ {
		angleByte[i] = ' '
	}
	fmt.Println(string(angleByte))

}

//拼接字符串
func TestAddString(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(s1)
	stringBuilder.WriteString(s2)
	fmt.Println(stringBuilder.String())
}

//定义枚举常量
func TestEnumConst(t *testing.T) {
	const (
		i = iota
		j
	)

	fmt.Println(i)
	fmt.Println(j)
}
