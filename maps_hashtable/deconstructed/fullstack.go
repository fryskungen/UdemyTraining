package main

import (
	"bufio"
	"fmt"
	"net/http"

	"log"
)

func main() {
	//get book
	result, err := http.Get("http://www.gutenberg.org/cache/epub/16712/pg16712.txt")
	if err != nil {
		log.Fatal(err)
	}
	//scan the oage
	scanner := bufio.NewScanner(result.Body)
	defer result.Body.Close()
	//set split
	scanner.Split(bufio.ScanWords)
	//create slices of string to hold the words
	buckets := make([][]string, 12)
	//create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	//loop over the words
	for scanner.Scan() {
		word := (scanner.Text())
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	fmt.Println(buckets[0:3])
	//print the length  of each bucket
	for i := 0; i < 12; i++ {

	//	fmt.Println(i, "--", len(buckets[i]))
	}
	//print the words in one of the buckets
	//fmt.Println(buckets[3])

}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
