// 修改 echo 程序输出 os.Args[0], 即命令的名字
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
