package main

import (
	"fmt"
)

func main() {
	a := [5]int{}
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	for v := range a {
		fmt.Println(v)
	}

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Printf("%v  \t %T \t %v \n", v, v, i)
	}
	/*fmt.Println(x[1:6])
	fmt.Println(x[6:])
	fmt.Println(x[3:8])
	fmt.Println(x[2:7])
	*/

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	fmt.Println(x)

	x1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 = append(x1[:3], x1[6:]...)
	fmt.Println(x1)

	fmt.Println("-------provinias-------")
	p := make([]string, 0, 7)
	p = append(p, "SJ", "H", "A", "C", "L", "P", "G")
	fmt.Println(len(p))
	fmt.Println(cap(p))

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i], i)
	}

	fmt.Println("-------Multi dim-------")
	m1 := []string{"James", "Bond", "Shaken,not stirred"}
	m2 := []string{"Miss", "Moneypenny", "I'm 008"}
	m := [][]string{m1, m2}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m2); j++ {
			fmt.Println(m[i][j])
		}
	}

	for i, v := range m {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}
