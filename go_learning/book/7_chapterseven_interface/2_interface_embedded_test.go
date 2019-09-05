package __chapterseven_interface

import (
	"io"
	"testing"
)

// 接口嵌套
// 只要实现了嵌套接口下的每个接口的方法，则这个接口中的所有嵌套接口的方法均可以被调用


//
/*type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}*/

type device struct{}

func (d *device) Write(p []byte) (b int, err error)  {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func TestEmbedded(t *testing.T)  {

	var wc io.WriteCloser = new(device)

	wc.Write(nil)

	wc.Close()


}
