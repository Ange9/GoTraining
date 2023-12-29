package main

import (
	"fmt"
	"math"
)

func main() {
	func() {
		fmt.Println("anon func")
	}()
	func(i int) {
		fmt.Println("anon func", 3)
	}(3)

	x := exp
	x()

	y := retFunc()
	fmt.Println(y())

	u := callBa(exp1)
	fmt.Println(u)

	p := powern(2)
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
}

func exp() {
	fmt.Println("Printing from expression")
}

func exp1() int {
	return 4
}

func retFunc() func() int {
	return func() int {
		return 3
	}
}

func callBa(f func() int) int {
	return f()
}

func powern(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
