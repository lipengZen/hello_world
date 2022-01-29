package main

// package helloworld

import "fmt"

func three() (re [][]int) {
	fmt.Println(re)
	return
}

func main() {
	// func helloworld() {
	fmt.Println("hello, world")
	three()

	re2 := make([][3]int, 0)
	fmt.Println(re2)
}
