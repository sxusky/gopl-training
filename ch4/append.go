package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice 仍有增长空间，扩展slice内容
		z = x[:zlen]
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂度，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // 内置copy函数
	}
	z[len(x)] = y
	return z
}

func main() {
	var a []int
	b := 3
	fmt.Println(appendInt(a, b))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d \t %v\n", i, cap(y), y)
		x = y
	}
}
