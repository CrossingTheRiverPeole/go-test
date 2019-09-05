package __chapterfive_function

func Accumulate(value int)func()int  {
	//返回一个闭包
	return func() int{
		value++
		return value
	}
}


