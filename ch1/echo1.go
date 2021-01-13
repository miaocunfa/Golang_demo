package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//fmt.Printf("命令本身：%s", os.Args[0])
	fmt.Printf("%.2fs", time.Since(start).Seconds())
}
