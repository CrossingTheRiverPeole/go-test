package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T)  {

	m := map[int]func(op int)int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

// 使用map实现set： 添加元素、判断元素是否存在，删除元素、判断元素的个数
func TestMapForSet(t *testing.T)  {

	m := map[int]bool{}
	// 添加元素
	m[1] = true
	// 判断元素是否存在
	if m[3]{
		t.Log("3 is existing")
	}else{
		t.Log("3 is not existing")
	}
	//删除元素
	delete(m, 1)
	t.Logf("set的长度%d", len(m))
}
