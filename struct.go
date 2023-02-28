package main

import (
	"fmt"
	"sort"
)

func main() {
	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(q)

	q = q[1:5]
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b string
	}{
		{2, "10"},
		{3, "25"},
	}
	fmt.Println(s)

	arr := []int{50, 20, 10, 40, 30}
	sort.Ints(arr)
	fmt.Println(arr)

	var n = len(arr)
	fmt.Println(n)
	fmt.Println(cap(arr))

	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var myArr []int
	if myArr == nil {
		fmt.Println(myArr)
		fmt.Println("nil!")
	}
	myStruct := myVar{name: "John", age: 20}
	fmt.Println(myStruct)
	fmt.Println(myStruct.doSome())
}

type myVar struct {
	name string
	age  int
}

func (m *myVar) doSome() string {
	return m.name
}
