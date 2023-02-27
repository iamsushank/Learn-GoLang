package main

import (
	"fmt"
	"log"
)

type Customer struct {
	firstName string
	lastName  string
	Balance   float64
}

func main() {
	str := "Hello World"
	log.Println("String Str is assigned as", str)

	var a = 20
	fmt.Printf("Type: %T Value: %v\n", a, a)

	var b = "Sushank"
	fmt.Printf("Type: %T Value: %v\n", b, b)

	var c = 35.3
	fmt.Printf("Type: %T Value: %v\n", c, c)

	var d = 'a'
	fmt.Printf("Type: %T Value: %v\n", d, d)

	var e = true
	fmt.Printf("Type: %T Value: %v\n", e, e)

	//Declaring a struct
	var f = Customer{
		firstName: "Sushank",
		lastName:  "Mandal",
		Balance:   524874.254,
	}
	fmt.Printf("Type: %T Value: %v\n", f, f)

	customer := Customer{
		firstName: "Sushank",
		lastName:  "Mandal",
		Balance:   524874.254,
	}
	log.Println("customer struct:", customer)

	var g = 83 + 46i
	fmt.Printf("Type: %T Value: %v\n", g, g)

}

func learnPointer() {

}
