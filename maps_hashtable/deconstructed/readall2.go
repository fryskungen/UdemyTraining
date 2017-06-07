package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	result, err := http.Get("http://www.gutenberg.org/files/54734/54734-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(result.Body) //returns skuce of bytes BS - byteslice
	result.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(bs)))
	//set split funktion for scanning opertation
	scanner.Split(bufio.ScanWords)
	//print the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
