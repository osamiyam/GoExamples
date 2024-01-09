package main

import (
	"fmt"
	"math"
	"os"
)

var fname = "myfile.txt"

func work_write(n int) {
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
		fmt.Fprintf(fp, "%f %f %f\n", x, x * x, math.Sqrt(x))
	}
}

func work_read() {
	fp, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer func(){
		if err := fp.Close(); err != nil {
			panic(err)
		}
	}()
	for {
		var x1, x2, x3 float64
		_, err := fmt.Fscanln(fp, &x1, &x2, &x3)
		if err != nil{
			break
		}
		fmt.Printf("%f:%f:%f\n", x1, x2, x3)
	}
}

func main() {
	work_write(100)
	work_read()
	_ = os.Remove(fname)
}
	
