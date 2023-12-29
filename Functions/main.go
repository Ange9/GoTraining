package main

import "fmt"

func main() {
	foo()
	bar("Ange")
	s := aloha("Mocca")
	fmt.Println(s)

	s1, n := dogYears("Mocca", 40)
	fmt.Println(s1, n)

	y := sum(1, 2, 3, 4, 5)
	fmt.Println(y)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(p string) {
	fmt.Println("My name is", p)
}

func aloha(s string) string {
	return fmt.Sprint("They call me ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " age "), age
}

//variatic parameter
func sum(ii ...int) int {
	fmt.Println(ii)
	n := 0
	//is at the end a slice
	for _, v := range ii {
		n += v
	}
	return n
}

//defer: do not run until surrounding function finishes
