package main

import (
	"fmt"
)

var a int
var b int

func main() {

	ask()

	switch {

	case b > a:
		fmt.Println(b, "/", a, "=", int(b)/int(a), "with a remainder of", int(b)%int(a))
	case b < a:
		fmt.Println("Large number must be larger than small number, try again")
		ask()
	}

}

func ask() {
	fmt.Print("Enter one small number please: ")
	fmt.Scan(&a)
	fmt.Print("Enter one large number please: ")
	fmt.Scan(&b)
}
