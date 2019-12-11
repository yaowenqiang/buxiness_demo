package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("r.MatchString(\"peach\")", r.MatchString("peach"))

	fmt.Println("r.FindString(\"peach punch\")", r.FindString("peach punch"))

	fmt.Println("r.FindStringIndex(\"peach punch\")", r.FindStringIndex("peach punch"))

	fmt.Println("r.FindStringSubmatch(\"peach punch\")", r.FindStringSubmatch("peach punch"))
	fmt.Println("r.FindAllString(\"peach\", -1)", r.FindAllString("peach", -1))
	fmt.Println("r.FindAllStringSubmatchIndex(\"peach punch pinch\", -1)", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println("r.Match([]byte(\"peach\")", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println("r.ReplaceAllString(\"a peach\", \"<fruit>\")", r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
