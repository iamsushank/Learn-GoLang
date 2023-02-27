package main

import (
	"fmt"
)

func main() {

	var str string = "Hello \nworld"
	fmt.Println(str)

	var num int = 15
	fmt.Println(num)

	fmt.Println("Hello, World!")
	fmt.Println(doSumOfTwoNumber(10, 20))

	fmt.Print(doProductAndDivide(45.5, 32.33))

}

func doProductAndDivide(a float32, b float32) (float32, float32) {
	return a * b, a / b
}

func doSumOfTwoNumber(a int, b int) int {
	return a + b
}
