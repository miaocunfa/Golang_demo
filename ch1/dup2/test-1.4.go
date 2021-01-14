package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(arg, f, counts)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
			//fmt.Printf("%s\t%d\n", line, n)
		}
		//fmt.Printf("%d\t%s\n", n, line)
	}
}

//countLines 增加文件名参数
func countLines(arg string, f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		filenameContent := arg + ": " + input.Text()
		counts[filenameContent]++
	}
}