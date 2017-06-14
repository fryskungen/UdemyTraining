package main

import (
	"fmt"
)

var antal = 100

func main() {

	in := generate()
	s := factorial(in)

	for sum := range s {
		fmt.Println(sum)
	}

}
func generate() chan float64 {		//using type int made 100 facts impossible
	out := make(chan float64)
	go func() {
		for i := 0; i <= antal; i++ {
			out <- float64(i)
		}
		close(out)
	}()
	return out

}

func factorial(in chan float64) chan float64 {
	out := make(chan float64)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n float64) float64 {

	var total float64 = 1
	for i := n; i > 0; i-- {
		total *= i

	}
	return total
}
