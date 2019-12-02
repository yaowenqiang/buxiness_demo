package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("1234.5", 0, 64)
	fmt.Println(i)
}
