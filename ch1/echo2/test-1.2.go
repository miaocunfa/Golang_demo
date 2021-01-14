package main

import (
	//"fmt"
	"os"
)

func main() {
	//s, sep := "", ""
	for index, arg := range os.Args[1:] {
		println(index, arg)
	}
}
