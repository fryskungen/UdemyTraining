package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter one small number please: ")
	fmt.Scan(&a)
	fmt.Print("Enter one large number please: ")
	fmt.Scan(&b)

	if b < a {
		fmt.Println("Large number must be larger than small number, try again")

	} else {
		fmt.Println(int(b)/int(a), "with a remainder of", int(b)%int(a))
	}

}
// ÖVERKURS! ställ frågan igen vid b < a