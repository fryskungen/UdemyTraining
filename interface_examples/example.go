package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}
type Square struct {
	side float64
}

type Triangle struct {
	side float64
}
type Circle struct {
	radius float64
}

func (z Triangle) area() float64 {
	return (z.side * z.side) / 2
}

func (z Square) area() float64 {
	return z.side * z.side
}
func (b Circle) area() float64 {
	return math.Pi * b.radius * b.radius

}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	t := Triangle{5}
	c := Circle{10}
	info(s)
	info(t)
	info(c)

}
