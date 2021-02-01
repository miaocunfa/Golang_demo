package main

import (
	"../tempconv"
	"fmt"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

	fmt.Printf("0K = %s\n", tempconv.KToC(0))
	fmt.Printf("1K = %s\n", tempconv.KToC(1))
}
