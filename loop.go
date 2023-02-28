package main

import (
	"fmt"
	"math"
)

type obj struct {
	roll  int
	name  string
	marks float64
}

func main() {

	student1 := obj{
		roll:  25,
		name:  "Sushank",
		marks: 53.25,
	}
	student2 := obj{
		roll:  25,
		name:  "Sushank",
		marks: 53.25,
	}
	var slice []obj
	slice = append(slice, student1)
	slice = append(slice, student2)

	for i, k := range slice {
		fmt.Println(i, k.name, k.marks, k.roll)
		fmt.Println("///////////////////")
	}

	i := 1
	for i < 1000 {
		i += i
	}
	fmt.Println(i)

	fmt.Println(sqrt(2), sqrt(-4))

	if i := 10; i < 20 {
		fmt.Println(i)
	}

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	var i, j = 5, 20

	if i < j {
		fmt.Println(j)
	}
	return fmt.Sprint(math.Sqrt(x))
}
