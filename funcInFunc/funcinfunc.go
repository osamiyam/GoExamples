package main

import "fmt"

func abc(i int) int {
	var xyz = func(j int) int {
		var stu = func(k int) int {
			return k * 3
		}
		return j * 2 * stu(j)
	}
	return i * xyz(i)
}

func mymap(f func(int) int, lst []int) []int {
	lst1 := []int{}
	for i := 0; i < len(lst); i++ {
		lst1 = append(lst1, f(lst[i]))
	}
	return lst1
}

func make_obj(x, y int) func (string, ...int) int {
	var foo = func (msg string, args ...int) int {
		switch msg {
		case "x":
			return x
		case "y":
			return y
		case "setx":
			x = args[0]
			return x
		case "sety":
			y = args[0]
			return y
		case "print":
			fmt.Printf("(%d, %d)", x, y)
			return 0
		default:
			return -1
		}
	}
	return foo
}

func main() {
	fmt.Println(abc(20))
	var lst = []int{1, 2, 3, 4, 5}
	fmt.Println(mymap(func(x int) int { return x * x }, lst))
	objx := make_obj(3, 4)
	fmt.Println(objx("x"), objx("y"))
	objx("setx", 23)
	objx("sety", 55)
	objx("print"); fmt.Println()
	fmt.Println(objx("x"), objx("y"))
}
