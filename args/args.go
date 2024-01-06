
package main

import (
  "fmt"
  "os"
)

func main() {
	m := os.Args
	fmt.Println(m)
	fmt.Println(len(m))
}
