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
	res1, res2 := <-ch, <-ch
	fmt.Printf("fib(%d) = %d\n", fibN, res1 + res2)
}

func work2() {
	fmt.Printf("fib(%d) = %d\n", fibN, fib(fibN))
}

func work3() {
	ch := make(chan int)
	go fibwork(fibN - 2, ch)
	res := fib(fibN - 1)
	fmt.Printf("fib(%d) = %d\n", fibN, res + <-ch) 
}

func main() {
	fmt.Printf("\n** Computing fib(%d) **\n\n", fibN)
	start := time.Now()
	work1()
	fmt.Println("Two threads & main: ", time.Since(start))
	start = time.Now()
	work2()
	fmt.Println("Single thread: ", time.Since(start))
	start = time.Now()
	work3()
	fmt.Println("One thread & main: ", time.Since(start))
	fmt.Println()
}
