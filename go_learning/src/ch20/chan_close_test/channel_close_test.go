package chan_close_test

import (
	"fmt"
	"sync"
	"testing"
)

/**
生产者：产生数据，并放到channel中
 */
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

/**
消费者：获取生产者产生的数据
 */
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for{
			if data, ok := <-ch; ok {
				fmt.Println("获取到数据", data)
			}else{
				break
			}
		}
		wg.Done()
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
