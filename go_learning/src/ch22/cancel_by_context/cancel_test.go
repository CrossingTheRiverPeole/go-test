package cancel_by_context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			for {
				if isCancel(ctx) {
					break
				}
			}
			fmt.Println(i, "calceld")
		}(ctx)
		cancel()
		time.Sleep(time.Second * 1)
	}
}

/**
测试使用context取消任务，注：如果父context取消了，子context会级联取消
 */

func CancelTest(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestContesxtCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			for {
				if CancelTest(ctx) {
					break
				}
			}
			fmt.Println(i, "取消了任务")
			time.Sleep(1 * time.Second)
			wg.Done()
		}(ctx, i)
	}
	cancel()
	wg.Wait()
}
