package __chapterseven_interface

import (
	"fmt"
	"sort"
	"testing"
)

//接口定义的方法不要太多，一个接口描述它的功能即可，通过接口组合和嵌入扩展为复杂的接口

// 定义接口
type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("模拟写入数据")
	return nil
}

func TestInterfaceImpl(t *testing.T) {
	f := new(file)
	//声明一个DataWriter的接口
	var writer DataWriter
	//接口赋值
	writer = f
	//使用DataWriter接口进行数据写入
	writer.WriteData("data")
}

//-----------------多结构体实现同一个接口---------------------
type Service interface {
	Start()
	Log(string)
}

//
type Logger struct {
}

//Logger结构体实现Log方法
func (g *Logger) Log(l string) {
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {

}

//使用接口进行数据的排序
//将[]string定义为MyStringList类型
type MyStringList []string

//实现sort.interface获取元素数量的方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.interface 接口的比较元素的方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

//实现sort.interface接口的交换元素的方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func TestSortInterface(t *testing.T) {
	names := MyStringList{
		"3. Trip Kill",
		"5. Penta kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Kill",
	}

	//使用sort.interface包进行排序
	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}

//使用sort对常见类型进行排序
//对字符串切片进行排序
func TestSortSliceString(t *testing.T) {
	names := sort.StringSlice{
		"3. Trip Kill",
		"5. Penta kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Kill",
	}
	//也可以使用sort.Strings()进行排序

	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}

//对整型切片进行排序
func TestSortSliceInt(t *testing.T) {
	i := []int{
		9, 10, 4, 1, 2, 3,
	}
	sort.Ints(i)
	for _, v := range i {
		fmt.Printf("%d\n", v)
	}
}

//对结构体进行排序
type HeroKind int

//定义HeroKind常量，类似枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

//定义英雄的结构
type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

// 返回长度
func (h Heros) Len() int {
	return len(h)
}

//
func (h Heros) Less(i, j int) bool {
	if h[i].Kind != h[j].Kind {
		return h[i].Kind < h[j].Kind
	}
	//默认按英雄的名字字符升序排列
	return h[i].Name < h[j].Name
}

func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func TestSortStruct(t *testing.T) {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
		&Hero{"妲己", Mage},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v", v)
	}
}

func TestSortSlice(t *testing.T) {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
		&Hero{"妲己", Mage},
	}

	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind != heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})

	for _, v := range heros {
		fmt.Printf("%+v", v)
	}
}



