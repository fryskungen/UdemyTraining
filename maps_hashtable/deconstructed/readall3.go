package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//get book
	result, err := http.Get("http://www.gutenberg.org/files/8675/8675-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the oage
	//NewScanner takes a reader and result.Body implements the reader interface
	scanner := bufio.NewScanner(result.Body)
	defer result.Body.Close()
	//set the split funktion for the scanning operation
	scanner.Split(bufio.ScanWords)
	//loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
