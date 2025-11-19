package task_two

import "sync"

func generatePrint(num int, size int) {
	var wg sync.WaitGroup
	//声明一个channel
	ch := make(chan int, size) //声明size 即缓冲区大小，不声明缓冲区，默认无缓冲区即一发一收

	wg.Add(2)
	//声明一个协程 遍历数字 并发送到通道中
	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			ch <- i
			println("发送", i)
		}
		close(ch) //关闭通道 这里理解只有负责写入的协程需要关闭通道，只读的协程不需要关闭通道，否则会panic: close of closed channel
	}()

	//在声明一个协程 接收数据并打印
	go func() {
		defer wg.Done()
		for i := range ch {
			println("接收", i)
		}

	}()
	wg.Wait()

}
