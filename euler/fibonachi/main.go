package main

import "fmt"

func main() {
	var fiboeven = 0
	for i, j := 0, 1; j < 4000000; i, j = i+j,i {
		if i%2 == 0{
			fiboeven += i


		}

	}

	fmt.Println(fiboeven)
}