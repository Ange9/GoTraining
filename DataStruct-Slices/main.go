package main

import "fmt"

func main() {
	xs := []string{"one", "two", "three", "four"}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println(i, v)
	}

	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Println("------------------")
	fmt.Println(len(xs))

	fmt.Println(xs[0])

	fmt.Println("------------------")
	xs = append(xs, "five", "six")
	fmt.Println(xs)

	fmt.Println("------slice------------")
	fmt.Println(xs[2:4])
	fmt.Println(xs[:4])
	fmt.Println(xs[2:])

	fmt.Println("------delete------------")
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(xi)
	// delete #4
	xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)

	fmt.Println("------Make------------")
	x1 := make([]int, 0, 10)
	fmt.Println(cap(x1))
	fmt.Println(len(x1))

	fmt.Println("------Multi dim array------------")
	v1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	v2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	vs := [][]int{v1, v2}
	fmt.Println(vs)

	fmt.Println("------Copy slice------------")
	a1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	a2 := a1

	fmt.Println(a1)
	fmt.Println(a2)

	a1[0] = 7
	// both get changed, because both are pointing to the same memory address
	fmt.Println(a1)
	fmt.Println(a2)

	// instead use copy func
	a3 := make([]int, 6)
	copy(a3, a1)
	a1[1] = 7

	fmt.Println(a1)
	fmt.Println(a3)

}
