package main

func main() {
	doMath(5, 2, add)
	doMath(5, 2, sub)

}

// Add function for ints.
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
