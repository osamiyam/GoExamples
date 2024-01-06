
package main

import "fmt"

func abc(i int) int {
	var xyz = func (j int) int {
		return j * 2
	}
	return i * xyz(i)
}

func main() {
	fmt.Println(abc(20))
}


