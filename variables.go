package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

type Customer struct {
	firstName string
	lastName  string
	Balance   float64
	timestamp time.Time
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
		timestamp: time.Now(),
	}
	log.Println("customer struct:", customer.timestamp)

	var g = 83 + 46i
	fmt.Printf("Type: %T Value: %v\n", g, g)

	var color = "Red"
	fmt.Println("value of str", color)
	learnPointer(&color)
	fmt.Println("value of str after learn pointer function", color)

}

func learnPointer(color *string) {
	*color = "Blue"

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

}
