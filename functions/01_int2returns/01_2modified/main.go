package main // modify "dualreturn" to a func expression. (minska scope, få allt i main. flyttade dualreturn in i main och := funcade den
		//assignt variable dualreturn as func
import "fmt"

var a = 4

func main() {
	dualreturn := func(n int) (int, bool, string, int) {
		return n / 2, (n%2 == 0), "remainder of", n % 2

	}
	fmt.Println(dualreturn(a))
}
