package main

import (
	"fmt"
)

func main() {
	s := factorial(4)
	for sum := range s {
		fmt.Println(sum)
	}

}

func factorial(n int) chan int {
	one := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i

		}
		one <- total
		close(one)
	}()
	return one
}
