package __chapterseven_interface

type State interface {
	//获取名称
	Name() string

	//改状态是否允许同状态转换
	EnableSameTransit() bool

	//响应状态开始时
	OnBegin()

	//响应状态结束时
	OnEnd()

	//判断能否

}

