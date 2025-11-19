package task_two

import (
	"sync"
	"sync/atomic"
)

func counter(count int) int {
	var wg sync.WaitGroup
	var sum int
	var mx sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mx.Lock()
			for i := 0; i < count; i++ {
				sum += 1
			}
			mx.Unlock()
		}()
	}
	wg.Wait()
	return sum
}

func counterAtomic(count int) int {
	var wg sync.WaitGroup
	var sum int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < count; i++ {
				atomic.AddInt32(&sum, 1)
			}
		}()
	}
	wg.Wait()
	return int(sum)
}
