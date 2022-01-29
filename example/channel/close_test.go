package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
}

func TestClose(t *testing.T) {

	ch := make(chan int, 5)

	close(ch)

	// i, ok := <-ch
	i := <-ch

	fmt.Println(i)

	var listChan chan *ListNode = make(chan *ListNode, 1)

	fmt.Println(listChan)
}
