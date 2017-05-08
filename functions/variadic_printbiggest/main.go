package main

import (
	"fmt"
)

func biggest(nummer ...int) int {
	var stor int
	for _, v := range nummer {
		if v > stor {
			stor = v // stor sätts till 0, loopar över range of numbers och om v är större än t
			// idigare nummer blir det "stor" osv tils största är hittat          (video 82 excercise 2 todd
		}

	}

	return stor
}

func main() {

	fmt.Println(biggest(1, 5, 6, 7, 11, 22, 14, 200, 125, 66))

}
