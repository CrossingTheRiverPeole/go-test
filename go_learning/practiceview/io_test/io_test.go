package io_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

//ioutl.ReadFile 读取文件的内容
func TestIOutil(t *testing.T)  {
	data, err := ioutil.ReadFile("a.txt")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(data))
	}
}

//分多次去取，每次读取指定长度的文件
func TestReadFile(t *testing.T)  {
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}else {
		//读取固定大小的
		buffer := make([]byte, 4096)

		length, err := f.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("读取了", length, "字节")
		fmt.Println(string(buffer))
	}
}


