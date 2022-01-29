package main

import (
	"fmt"
	"testing"
)

func TestSelectWrite(t *testing.T) {

	ch := make(chan int)

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	ch <- 0
	// 	close(ch)
	// }()

	for {
		select {
		case i := <-ch:
			fmt.Println("读管道:", i)
			break
			// case ch <- 0:
			// 	fmt.Println("写管道,如果没再打印读,那就阻塞了")
			// default:
			// 	fmt.Println("执行default,没有阻塞")
			// 	time.Sleep(time.Second)
		}
	}

}
