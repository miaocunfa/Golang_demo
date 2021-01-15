package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[1], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[0:1])
	fmt.Println(s[2:3])
	fmt.Println(s[0:12])
}
