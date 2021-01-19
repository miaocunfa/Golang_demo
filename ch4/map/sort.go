package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"mike": 35,
		"charlie": 34,
	}

	//var names []string
	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
