package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func is_leapyear(y int) bool {
	switch {
	case y%400 == 0:
		return true
	case y%100 == 0:
		return false
	default:
		return (y%4 == 0)
	}
}

func num_of_days(y int) int {
	s, i := 0, 1
	for i < y {
		s += 365
		if is_leapyear(i) {
			s += 1
		}
		i++
	}
	return s
}

func day_of_week(y int) int {
	return (num_of_days(y) + 1) % 7
}

var (
	nmon = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	header = "Su Mo Tu We Th Fr Sa"
	mnames = []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November",
		"December"}
)

func prcal(y int, m int) {
	ss := 0
	for i := 1; i < m; i++ {
		ss += nmon[i-1]
	}
	if is_leapyear(y) && m >= 3 {
		ss += 1
	}
	d, w, s := 1, (day_of_week(y)+ss)%7, ""
	days := nmon[m-1]
	if is_leapyear(y) && m == 2 {
		days += 1
	}
	fmt.Println()
	title := fmt.Sprintf("%s %d", mnames[m-1], y)
	for i := 0; i < (20-len(title))/2; i++ {
		fmt.Print(" ")
	}
	fmt.Println(title)
	fmt.Println(header)
	for i := 0; i < w; i++ {
		s += "   "
	}
	for d <= days {
		s += fmt.Sprintf("%2d ", d)
		d += 1
		w = (w + 1) % 7
		if w == 0 {
			fmt.Println(s)
			s = ""
		}
	}
	fmt.Println(s)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		date := time.Now()
		prcal(date.Year(), int(date.Month()))
	} else if len(args) == 2 {
		m, _ := strconv.Atoi(args[0])
		y, _ := strconv.Atoi(args[1])
		prcal(y, m)
	}
}
