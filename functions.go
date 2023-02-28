package main

import "fmt"

func main() {
	myFun := func(a, b int) int {
		return a + b
	}
	println(myFun(20, 20))
	fmt.Println(nestedFunc(myFun))
}

func nestedFunc(myFun func(int, int) int) int {
	return myFun(10, 20)
}
