package task_two

import (
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup //这是一个协程计数器

	numArr := getNumArr(true, 1, 100)
	numArr2 := getNumArr(false, 1, 100)
	wg.Add(2)

	go func() {
		defer wg.Done()
		concurrentPrint(numArr)

	}()

	go func() {
		defer wg.Done()
		concurrentPrint(numArr2)

	}()
	wg.Wait() //等待所有协程执行完毕，有点类似java的CompletableFuture.allOf()
}

func TestTasks(t *testing.T) {
	runTasks(createTasks(5))
}
