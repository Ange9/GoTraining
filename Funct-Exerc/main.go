package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Printf("Name %v, age %v", p.first, p.age)
}

func main() {
	//defer fmt.Println("defered")
	fmt.Println(foo())
	fmt.Println(bar())
	//defer fmt.Println("defered 2")

	fmt.Println("--------Variadic func------------")
	v := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo1(v...))
	fmt.Println(bar1(v))

	fmt.Println("--------Method------------")
	p := Person{
		first: "Ange",
		age:   30,
	}
	p.Speak()

}
func foo() int {
	return 2
}
func bar() (int, string) {
	return 4, "Age"
}

func foo1(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar1(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
