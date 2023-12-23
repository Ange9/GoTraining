package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v \t", x)

	if x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if 100 < x && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else {
		fmt.Println("Between 201 and 250")
	}

	fmt.Printf("The value is %v \t \n", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))
	// fmt.Printf("The value is %v \t", rand.Intn(3))

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100")
	case 100 < x && x <= 200:
		fmt.Println("Between 101 and 200")
	default:
		fmt.Println("Between 201 and 250")
	}
}
