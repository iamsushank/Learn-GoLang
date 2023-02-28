package main

import "fmt"

func main() {

	myVar := "dog"
	switch myVar {
	case "cat":
		fmt.Println("this is cat")

	case "dog":
		fmt.Println("this is dog")

	default:
		fmt.Println("don't know what it is")
	}

	a := 10
	if a < 2 {

	} else if a < 20 {
		println("10 is less than 20")
	}

}
