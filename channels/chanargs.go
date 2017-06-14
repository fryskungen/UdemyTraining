package main

import (
	"fmt"
)

func main() {
	c := incrementor(4)
	cSum := puller(c)
	fmt.Println(<-cSum)
}

func incrementor(n int) chan int {
	one := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			one <- i
		}
		close(one)
	}()
	return one
}

func puller(c chan int) chan int {
	two := make(chan int)
	go func() {
		var sum = 1
		for n := range c {
			sum *= n
		}
		two <- sum
		close(two)
	}()
	return two
}
