package main

import "fmt"

func kuku(n int) {
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := i * j
			fmt.Printf("%5d", k)
		}
		fmt.Println()
	}
}

func main() {kuku(9)}
