package main

import "fmt"

func foo(numbers ...int)  {
	fmt.Println(numbers)
}

var aSlice = []int{1, 2, 3, 4}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	foo(aSlice...)
	foo()

}
