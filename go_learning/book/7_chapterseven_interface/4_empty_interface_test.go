package __chapterseven_interface

import (
	"fmt"
	"testing"
)

//空接口是接口类型的特殊形式，空接口没有任何方法，可以接收任何方法

//空接口接收值，然后再转换为别的值

func TestEmptyInterface(t *testing.T) {
	a := 1
	var b interface{}
	b = a
	var c int
	c = b.(int)
	fmt.Println(c)
}

//示例：使用空接口实现可以保存任意值的字典

type Dictionary struct {
	data map[interface{}]interface{}
}

//根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

//设置键值
func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

//遍历所有的字段，如果回调函数返回false，停止遍历
func (d *Dictionary) Visit(callback func(key, value interface{}) bool) {
	if callback == nil {
		return
	}
	for key, value := range d.data {
		if !callback(key, value) {
			return
		}
	}
}

//清空所有资源
func (d *Dictionary) clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.clear()
	return d
}

func TestDictionary(t *testing.T) {

	//初始化字典
	d := NewDictionary()

	//添加数据
	d.Set("My factory", 60)
	d.Set("Terra Craft", 36)
	d.Set("Don't hungry", 24)

	//获取值并打印数据
	fmt.Println(d.Get("Terra Craft"))

	//遍历所有的元素
	d.Visit(func(key, value interface{}) bool {
		if value.(int) > 30 {
			fmt.Println(key, "is expecsive")
			return true
		}
		fmt.Println(key, "is cheap")
		return false
	})
}
