package main

import (
	"fmt"
	"math"
	"os"
)

func work(n int) {
	fname := "myfile.txt"
	fp, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			panic(err)
		}
	}()
	for i := 0; i < n; i+= 1 {
		x := float64(i)
		fmt.Fprintf(fp, "%f, %f, %f\n", x, x * x, math.Sqrt(x))
	}
}

func main() {
	work(100)
}
	
