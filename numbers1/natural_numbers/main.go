package main

import "fmt"

func main() {
	var sum1 = 0
	var sum2 = 0
	for i := 0; i < 1000; i++ {

		switch {
		case i%3 == 0:
			sum1 += i
		case i%5 == 0:
			sum2 += i
		}

	}

	fmt.Println(sum1 + sum2, "BOOOYAH!")
}
