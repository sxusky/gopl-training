package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1:%x\nc2:%x\n%t\nc1:%T c2:%T", c1, c2, c1 == c2, c1, c2)
}
