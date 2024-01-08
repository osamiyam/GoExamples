package main

import(
	"fmt"
	"test/vec2d"
)


func main() {
	v1 := vec2d.Vec{2, 3}
	v2 := vec2d.Vec{5, 6}
	fmt.Println(v1, v2)
	fmt.Println(v1.Length())
	fmt.Println(v1.Add(v2))
	fmt.Println(v1.InnerProd(v2))
}
