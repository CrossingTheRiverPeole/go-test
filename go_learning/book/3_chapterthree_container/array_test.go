package __chapterthree_container

import (
	"fmt"
	"testing"
)

//数组定义
func TestArrayDefine(t *testing.T) {
	var team [3]string
	team[0] = "ss"
	team[1] = "sdd"
	fmt.Println(team)
	fmt.Println(len(team))
	//定义时即初始化数组
	team2 := [2]string{"f"}
	fmt.Println(team2)

	// 遍历数组
	for k, v := range team {
		fmt.Println(k)
		fmt.Println(v)
	}

}

//切片--没有固定大小，第一次分配时是1，每次扩展两倍；切片包含地址、容量、大小,切片即是一段连续的内存
func TestSlice(t *testing.T) {
	//定义切片
	a := [4]int{1, 2}
	fmt.Println(a[3:4])

	//定义切片
	b := []int{1}
	fmt.Println(b)

	//重置切片
	c := b[0:0]
	fmt.Println(c)

	//声明切片，此时切片为nil
	var d []int //
	fmt.Println("append函数", append(d, 1))
	fmt.Println(d) // 此时d切片不会发生任何变化，append添加元素会重新生成一个新的地址的切片
	// 测试
	var test [5]int
	dd := test[2:]
	fmt.Println("dd", dd)
	fmt.Println("dd append", append(dd, 1))
	fmt.Println("dd", dd)


	// 声明并初始化切片
	e := []int{} //此时切片已经在内存中分配地址，但是切片为空
	fmt.Println("eee", e)

}

//向切片中添加元素:使用append向切片中添加元素时，这是会重新生成一个切片
func TestAppendSlice(t *testing.T)  {
	var number []int
	fmt.Printf("%p\n", number)
	number = append(number,1,2,3)
	//fmt.Println(append(number, 1,2,3))
	fmt.Println(append(number,4))
	fmt.Println(number)

}

func TestMakeSlice(t *testing.T)  {
	a := make([]int,2,5)
	fmt.Println(a)
}
//删除切片中的数据
func TestSliceData(t *testing.T)  {
	// 定会切片
	seq := []int{1,2,3}
	//删除指定位置的元素
	index := 1
	seq = append(seq[:index],seq[index +1:]...)
	fmt.Println(seq)
}
