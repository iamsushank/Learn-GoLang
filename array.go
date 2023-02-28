package main

import "fmt"

func main() {

	arr := make([]int, 5)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr, len(arr), cap(arr))

	b := arr[:3]
	fmt.Println(b, len(b), cap(b))

	a := b[1:5]
	fmt.Println(a, len(a), cap(a))

	c := a[1:3]
	fmt.Println(c, len(c), cap(c))

	c = append(c, 6, 7, 8)
	fmt.Println(c, len(c), cap(c))

	for i, v := range c {
		fmt.Println(i, v)
	}

	for i := range c {
		c[i] = i * 2
	}

	for _, v := range c {
		fmt.Println(v)
	}

}
