package main

import "fmt"

// pc[i] 是i的种群统计
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Printf("%d 的种群统计为 %d", 6, PopCount(6))
}

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])

}
