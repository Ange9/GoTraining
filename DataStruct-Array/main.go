package main

import "fmt"

func main() {
	//array: values of the same type, dont change in size, not very used in coding
	a := [3]int{41, 42, 43}
	fmt.Println(a)

	b := [...]string{"Heello", "there"}
	fmt.Println(b)

	var c [2]int
	c[0] = 2
	c[1] = 8

	fmt.Println(c)
	fmt.Println(len(c))

	//slice: values of the same type, changes in size, are built on arrays

	//map: key value storage, unorder group of values of one type

	//struct: composite/aggregate, collects values of different types together

	//scope, used to collapse a block of code, variables declared inside the scope are only accessible within the scope
	{
		x := 3
		fmt.Println(x)
	}

}
