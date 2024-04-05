package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	names := []string{"a", "A"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
	// a := []byte("a")
	// A := []byte("A")
	// [a, A] -> 97 65
	// fmt.Printf("a: %v, A: %v\n", a[0], A[0])
	// fmt.Println("a" > "A")
}