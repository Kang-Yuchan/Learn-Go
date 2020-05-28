package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
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
