package main

import "fmt"

func main() {

	var i int = 42
	var j int = 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	println("/////////////////////////////////")

	var a int = 15

	x := &a
	fmt.Println(*x)

	*x = 50
	fmt.Println(a)

}
