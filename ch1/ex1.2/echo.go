// 修改 echo 程序， 输出参数的索引和值，每一行
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, ":", arg)
	}
}
