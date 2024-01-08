package main

import (
  "fmt"
  "os"
  "strconv"
)

func fib(n int) int {
  if n <= 1 {
    return 1
  } else {
    return fib(n - 1) + fib(n - 2)
  }
}

func main() {
  args := os.Args
  if len(args) < 2 {
     fmt.Println("Usage: fib n")
  } else {
  	m, err := strconv.Atoi(os.Args[1])
  	if err != nil {
		fmt.Println("Error: ", err)
	} else {
  		for i := 0; i <= m; i++ {
  			fmt.Println(i, fib(i))
  		}
	}
  }
}
