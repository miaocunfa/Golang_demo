package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("命令是： %s\n", os.Args[0])
	fmt.Printf("参数是： %s\n", s)
}
