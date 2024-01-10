package main

import (
	"fmt"
	"time"
)

const fibN = 38

func fibwork(n int, ch chan int) {
	ch <- fib(n)
}

func fib(n int) int {
	if (n <= 1) {
		return 1
	} else {
		return fib(n - 1) + fib(n - 2)
	}
}

func work1() {
	ch := make(chan int)
	go fibwork(fibN - 2, ch)
	go fibwork(fibN - 1, ch)
	fmt.Printf("fib(%d) = %d\n", fibN, <-ch + <-ch) 
}

func work2() {
	fmt.Printf("fib(%d) = %d\n", fibN, fib(fibN))
}

func main() {
	fmt.Printf("\n** Computing fib(%d) **\n\n", fibN)
	start := time.Now()
	work1()
	fmt.Println("Two threads: ", time.Since(start))
	start = time.Now()
	work2()
	fmt.Println("Single thread: ", time.Since(start))
	fmt.Println()
}
