package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	fmt.Printf("%s", bs)
}
