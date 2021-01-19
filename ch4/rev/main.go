package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]
}

// 就地翻转一个整型 slice 中的元素
func reverse(s []int)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}