// Nonempty 演示了 slice 的就地修改算法
package main

import "fmt"

// nonempty 返回一个新的 slice, slice 中的元素都是非空字符串
// 在函数的调用过程中, 底层数组的元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // 引用原始 slice	的新的零长度的 slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return  out
}

func main() {
	data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data))  // ["one" "three"]
	//fmt.Printf("%q\n", data)            // ["one" "three" "three"]

	fmt.Printf("%q\n", nonempty2(data))  // ["one" "three"]
	fmt.Printf("%q\n", data)             // ["one" "three" "three"]
}
