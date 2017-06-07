package main

import (
	"fmt"
)

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	word := "hello"
	letter = rune(word[1])
	fmt.Println(letter)

	for i := 65; i <= 122; i++ {
		fmt.Println(i, " -- ", string(i), " -- ", i%12) //getting "buckets" to put words in
	}
}
