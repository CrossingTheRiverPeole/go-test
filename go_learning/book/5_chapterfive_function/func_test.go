package __chapterfive_function

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"testing"
)

// 函数的参数
func add(a, b int) int {
	return a + b
}

// 函数的返回值：多个返回值；多个相同类型返回值
func test(a, b int) (int, error) {
	err := errors.New("计算出错")
	return a + b, err
}

//将秒解析为时间单位
const (
	//定义每分钟的秒数
	SecondsPerMinute = 60
	//定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	//定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

func resolverTime(seconds int) (day int, hour int, minute int) {
	minute = seconds / SecondsPerMinute
	hour = seconds / SecondsPerHour
	day = seconds / SecondsPerDay
	return
}

func TestFun(t *testing.T) {
	fmt.Println(resolverTime(1000))
	//只获取小时和分钟
	_, hour, minute := resolverTime(18000)
	fmt.Println(hour, minute)

	//只获取天
	day, _, _ := resolverTime(90000)
	fmt.Println(day)
}

//参数传递：指针、切片、map等引用类型在传递时不会发生复制，而是将指针进行复制，其他的参数都是进行参数复制
func passByValue() {

}

func fire() {
	fmt.Println("fire")
}

//函数作为参数进行传递
func TestFuncByVaribale(t *testing.T) {
	var b func()
	b = fire
	// 调用函数
	b()
}

func visit(list []int, f func(int))  {
	for _, v := range list {
		f(v)
	}
}

// //匿名函数作为回调函数
func TestAnonymousFunc2(t *testing.T){
	visit([]int{1,2,3}, func(i int) {
		fmt.Println(i)
	})
}

//使用匿名函数实现封装操作
var skillParam = flag.String("skill", "", "skill to perform")

func TestAnonymousFunc3(t *testing.T)  {
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("fire chicked")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angle fly")
		},
	}

	for key, value := range skill  {
		fmt.Printf("%T %T", key, value)
	}

	if f, ok := skill[*skillParam]; ok{
		f()
	}else{
		fmt.Println("skill not found")
	}
}

//字符串链式处理
func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		// 第一个需要处理的字符串
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func TestFuncChain(t *testing.T) {
	list := []string{
		"go scanner",
		"go parse",
		"go compiler",
		"go printer",
		"go formater",
	}
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	//处理字符串
	StringProcess(list, chain)

	//输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}

//匿名函数

//定义时调用匿名函数
func TestAnonymousFunc(t *testing.T) {
	//调用匿名函数
	func(data int) {
		fmt.Println(data)
	}(100)
	// 将匿名函数赋值给变量
	f := func(data int) {
		fmt.Println(data)
	}
	f(100)

}









//函数定义：返回值参数以及
func TestDefineFunc(t *testing.T) {

}
