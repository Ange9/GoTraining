package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 2
	for i > 0 {
		// fmt.Println(i)
		i--
	}

	i = 2
	for i > 0 {
		// fmt.Println(i)
		if i == 5 {
			break
		}
		i--
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		// fmt.Println(i)
	}

	// xi := []int{42, 43, 44, 45, 46, 47}
	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }
	fmt.Println("-----------")

	m := map[string]int{
		"James": 42,
		"Money": 65,
	}
	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	fmt.Println("-----------")
	// comma ok
	age := m["James"]
	fmt.Println(age)

	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no value for Q", v)
	}
	fmt.Println("-----------")
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is 3")

		}
	}

}
