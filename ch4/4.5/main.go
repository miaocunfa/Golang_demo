package main

import "fmt"

func removeDuplicates(strings []string) []string {
	out := strings[:1]
	for i := 1; i < len(strings); i++ {
		if strings[i] != out[len(out)-1] {        // 判断当前字符 与栈顶字符是否相同
			out = append(out, strings[i])
		}
	}
	return out
}

func main() {
	s := []string{"a", "b", "b", "b", "c", "d", "d"}
	s = removeDuplicates(s)
	fmt.Println(s)
}
