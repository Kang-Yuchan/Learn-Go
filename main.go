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

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	myName := "Kang Yuchan"
	fmt.Println(myName, multiply(12, 30))
	fmt.Println(lenAndUpper(myName))
	repeatMe("yuchan", "miku", "masa", "junko")
	result := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(result)
	fmt.Println(canIDrink(12))
}
