package _map

import (
	"testing"
)

func TestDefine(t *testing.T)  {

	//三种定义方式，一般是使用第一种定义方式
	m1 := map[int]string{1:"name", 2:"age"}
	t.Log(m1)

	m2 := map[int]string{}
	m2[1] = "test"

	//使用make创建map
	m3 := make(map[int]int,10)
	t.Logf("m3 length %d",len(m3))
}

func TestAccessExisting(t *testing.T)  {
	//map定义之后如果没有进行初始化，value是值类型的默认值
	m1 := map[int]int{}
	t.Log(m1[1])

	//判断map中key是否存在
	if v,ok := m1[2];ok {
		t.Logf("m[2] value is %d", v)
	}else{
		t.Log("key is not existing")
	}

	// 遍历map
	for k,v := range m1{
		t.Logf("key: %d,value:%d", k,v)
	}
}
