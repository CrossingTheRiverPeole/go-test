package objpool_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/**
使用缓冲channel实现对象池
 */

type ReuseableObj struct {
	// 池中对象
}

//缓冲池
type ObjPool struct {
	buffchan chan *ReuseableObj
}

func NewObjPool(numObj int) *ObjPool {
	objPool := &ObjPool{}
	objPool.buffchan = make(chan *ReuseableObj, numObj)
	for i := 0; i < numObj; i++ {
		objPool.buffchan <- &ReuseableObj{}
	}
	return objPool
}

/**
获取池中的对象
 */
func (p *ObjPool) GetObjPool(timeout time.Duration) (*ReuseableObj, error) {
	select {
	case ret := <-p.buffchan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReuseableObj) error {
	select {
	case p.buffchan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i :=0; i< 11 ;i++  {
		if v,err := pool.GetObjPool(1 * time.Second); err != nil {
			t.Error(err)
		}else{
			fmt.Printf("%T\n", v)
			if e := pool.ReleaseObj(v);e != nil {
				t.Error(e)
			}
		}
	}
		fmt.Println("Done")

}
