package main

import (
	"fmt"
)

func main() {
	var a [3]int				// 3个整数的数组
	fmt.Println(a[0])			// 输出数组的第一个元素
	fmt.Println(a[len(a)-1])    // 输出数组的最后一个元素, 即[2]

	// 输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//q := [...]int{1, 2, 3}
	//fmt.Printf("%T\n", q) // [3]int

	//q := [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4} // cannot use [4]int literal (type [4]int) as type [3]int in assignment

	type Currency int
	const  (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(USD, symbol[USD]) // "0 $"
	fmt.Println(EUR, symbol[EUR]) // "1 €"
	fmt.Println(GBP, symbol[GBP]) // "2 £"
	fmt.Println(RMB, symbol[RMB]) // "3 ¥"
}
