// 尝试测量可能低效的程序和使用 strings.Join 的程序在执行实际上的差异
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo2()
}

func echo1() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("echo1: %fs\n", time.Since(start).Seconds())
}

func echo2() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("echo2: %fs\n", time.Since(start).Seconds())
}
