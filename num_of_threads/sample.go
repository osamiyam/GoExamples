
package main

import (
   "fmt"
   "runtime"
)

func mugen() {
	fmt.Println("A thread started")
	for {
	}
}

func main() {
	fmt.Println("# of CPUs = ", runtime.NumCPU())
	fmt.Println("# of current threads = ", runtime.NumGoroutine())
	go mugen()
	fmt.Println("# of current threads = ", runtime.NumGoroutine())
}
