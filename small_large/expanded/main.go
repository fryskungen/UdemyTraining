package main

import "fmt"

var a int
var b int

func main() {

	fmt.Print("Enter one small number please: ")
	fmt.Scan(&a)
	fmt.Print("Enter one large number please: ")
	fmt.Scan(&b)

	switch {

	case b > a:
		fmt.Println(int(b)/int(a), "with a remainder of", int(b)%int(a))
	default:
		fmt.Println("Large number must be larger than small number, try again")

	}

}
