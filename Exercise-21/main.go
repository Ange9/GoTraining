package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("The value of x is %v \n", x)
		fmt.Printf("The value of y is %v \n", y)

		if x < 4 && y < 4 {
			fmt.Println("Both are less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("Both are greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("X gte to 4 and lte to 6")
		} else if y != 5 {
			fmt.Println("y is not 5")
		} else {
			fmt.Println("none")
		}

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("Both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("X gte to 4 and lte to 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none")
		}
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

}
