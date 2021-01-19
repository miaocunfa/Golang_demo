package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var x1 []int
	x1 = append(x1, 1)
	x1 = append(x1, 2, 3)
	x1 = append(x1, 4, 5, 6)
	x1 = append(x1, x1...)   // 追加 x 中的所有元素
	fmt.Println(x1)          // [1 2 3 4 5 6 1 2 3 4 5 6]
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x){
		// slice 仍有增长空间, 扩展 slice 内容
		z = x[:zlen]
	} else {
		// slice 已无空间, 为它分配一个新的底层数组
		// 为了达到分摊线性复杂性, 容量扩展一倍
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 内置 copy 函数
	}
	z[len(x)] = y
	return z
}
