package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	myName := "Kang Yuchan"
	fmt.Println(myName, multiply(12, 30))
	fmt.Println(lenAndUpper(myName))
	repeatMe("yuchan", "miku", "masa", "junko")
}
