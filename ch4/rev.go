package main

import "fmt"

// 就地反转一个整型slice中的元素

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[1:])
	reverse(a[1:])
	fmt.Println(a)
}
