package main

import "fmt"

func main() {
	// 示例1
	//ages := make(map[string]int)

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	ages["alice"] = 32
	fmt.Println(ages["alice"])  // 32

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
