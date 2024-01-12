package main

import (
	"fmt"
	"sync"
)

const (
	N        = 100
	nthreads = 10
)

func producer(ch chan int) {
	for i := 1; i <= N; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, no int, wg *sync.WaitGroup) {
	defer wg.Done()
	data := []int{}
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("This is Process %d. I got %d.\n", no, i)
		data = append(data, i)
	}
	fmt.Printf("Process %d: All numbers I've recieved are %d.\n", no, data)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 0; i < nthreads; i++ {
		wg.Add(1)
		go consumer(ch, i, &wg)
	}
	producer(ch)
	wg.Wait()
}
