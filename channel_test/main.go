package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int64, 100)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(ch <-chan int64) {
		defer waitGroup.Done()
		for i := range ch {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(i)
		}
	}(ch)

	for i := int64(0); i < 100; i++ {
		ch <- i
	}

	close(ch)
	fmt.Println("close chan")
	waitGroup.Wait()
	fmt.Println("结束！")
}
