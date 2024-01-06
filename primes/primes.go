package main

import "fmt"

func primes(n int) {
	var a [] bool
	for i := 0; i < n; i++ {a = append(a, true)}
	for i := 2; i < n; i++ {
		if a[i] {
			fmt.Print(i, " ")
			for j := i * 2; j < n; j += i {
				a[j] = false
			}
		}
	}
	fmt.Println()
}

func main() {primes(1000)}


