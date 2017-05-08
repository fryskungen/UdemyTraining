package main

import "fmt"

var a = 4

func main() {
	fmt.Println(dualreturn(a))
}
func dualreturn(n int) (int, bool, string, int) { //dual went multi-return
	return n/2, (n % 2 ==0), "remainder of", n % 2

}
