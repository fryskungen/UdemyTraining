package main

import (
	"bufio"
	"fmt"

	"strings"
)

func main() {
	input := "Once upon a time a little mouse and a little sausage, who loved each other like sisters, decided to live together, and made their arrangements in such a way that every day one would go to walk in the fields, or make purchases in town, while the other remained at home to keep the house."
	scanner := bufio.NewScanner(strings.NewReader(input))
	//set split funktion for scanning opertation
	scanner.Split(bufio.ScanWords)
	//count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
