package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	s := square{
		length: 5,
		width:  5,
	}
	c := circle{
		radius: 5,
	}

	fmt.Printf("Circles area is: %v ", info(c))
	fmt.Printf("Square area is: %v ", info(s))
}
