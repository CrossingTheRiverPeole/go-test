package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T)  {

	//字符串分割与连接
	s := "A,B,C"
	splits := strings.Split(s, ",")
	t.Log(splits)

	//字符串连接
	join := strings.Join(splits, "_")
	t.Log(join)
}

/**
字符串类型转换：
 */
func TestStringConv(t *testing.T)  {
	//字符串转换为整型
	if i, e := strconv.Atoi("10"); e == nil{
		fmt.Println(10 + i)
	}

	// int转换为字符串
	s := strconv.Itoa(10)
	t.Log(s)
}
/**

 */
func TestRune(t *testing.T)  {
	s := "中华人民共和国"
    //sr := []rune(s)
	for _, v := range s{
		t.Logf("%c",v)
	}

	for i := 0; i < len(s); i++ {
		t.Logf("%c", s[i])
	}
}

