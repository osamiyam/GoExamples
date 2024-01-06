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
  m, err := strconv.Atoi(os.Args[1])
  if err != nil {
     fmt.Println("!", err)
  }
  for i := 0; i <= m; i++ {
  	fmt.Println(i, fib(i))
  }
}
