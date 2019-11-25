package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)
	ints := []int{3, 5, 6, 4, 23}

	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
}
