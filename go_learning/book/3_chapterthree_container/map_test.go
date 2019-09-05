package __chapterthree_container

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	//常见map
	m := make(map[string]int)
	m["age"] = 1
	//判断是否存在键值
	if _, ok := m["age"]; ok {
		fmt.Println("存在键值")
	}
}

// 遍历map:map遍历是无序的
func TestRangeMap(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3

	for key, value := range m {
		fmt.Printf("%d : %d\n", key, value)
	}
}

//删除map中的元素
func TestDelMapElemet(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

	//直接进行初始化
	m2 := map[int]int{
		1:2,
	}
	fmt.Println(m2)
}

func TestEmptyMap(t *testing.T)  {
	m := map[int]int{
		1:2,
	}
	fmt.Println(m)
	m = map[int]int{}
	fmt.Println(m)
}

//并发map
func TestSyncMap(t *testing.T)  {
	//直接声明即可使用
	var scene sync.Map
	//保存
	scene.Store(1,2)
	scene.Store("name", "ss")

	//取值load
	fmt.Println(scene.Load(1))
	//删除
	scene.Delete("1")
	//遍历
	var name interface{}
	var v interface{}
	scene.Range(func(key, value interface{}) bool {
		name = key
		v = value
		fmt.Println(key,value)
		return true
	})
	fmt.Println(name,v)
}


