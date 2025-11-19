package task_two

import (
	"fmt"
	"sync"
	"time"
)

func getNumArr(odd bool, begin int, end int) []int {
	var numArr []int
	for i := begin; i < end; i++ {
		if odd {
			if i%2 == 1 {
				numArr = append(numArr, i)
			}
		} else {
			if i%2 == 0 {
				numArr = append(numArr, i)
			}
		}
	}
	return numArr

}

func concurrentPrint(nums []int) {
	for _, num := range nums {
		println(num)
	}
}

func createTasks(count int) []func() {
	var tasks []func()
	for i := 0; i < count; i++ {
		f := func() {
			b := i%2 == 0
			numArr := getNumArr(b, 1, 10)

			//模拟耗时
			time.Sleep(time.Duration(1000 * i))
			fmt.Println(numArr)
		}
		tasks = append(tasks, f)
	}
	return tasks
}

// 设计一个函数，入参是一个函数集合，使用协程并发执行这些任务，同时统计每个任务的执行时间
func runTasks(tasks []func()) {
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			println(i, "任务执行时间：", i, end.Sub(start))
		}(task)
	}
	wg.Wait()
}
