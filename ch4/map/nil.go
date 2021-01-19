package main

import "fmt"

func main() {
	// 示例1 - nil
	var ages map[string]int
	fmt.Println(ages == nil)       // true
	fmt.Println(len(ages) == 0)    // true

	//ages["carol"] = 21             // panic: assignment to entry in nil map

	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("bob 不是字典中的值, ages == %d\n", age)
	}

	// 示例2 - equal
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42})) // 第一种: true, 第二种: false
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range  x {
		//if xv != y[k] {                             // 第一种
		if yv, ok := y[k]; !ok || yv != xv {          // 第二种
			return false
		}
	}
	return  true
}