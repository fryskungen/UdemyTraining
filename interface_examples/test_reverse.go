package main

import (
	"fmt"
	"sort"
)

type people []string

func (c people) Len() int { return len(c) }

func (c people) Less(i, j int) bool { return c[i] < c[j] }

func (c people) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func main() {
	studyGroup := people{"Zoe", "Benny", "Benson", "Frank", "Adrian", "Cedrian", "Dedrian"}
	s := []string{"Zeno", "John", "Al", "Jenny", "Forrest"}
	n := []int{1, 5, 6, 7, 9, 11, 2, 99, 14, 22}
	fmt.Println(studyGroup)

	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.StringSlice(s))) //IntSlice/StringSlice implementerar Interface interface genom metoderna de defienras med.
	fmt.Println(s)				     // "konvertera" s till en STringslice så funkar det

}
