package main

func main() {
	myMap := make(map[string]string)
	myMap["name"] = "Sushank"
	myMap["age"] = "24"
	println(myMap["name"])

	mp := make(map[int]string)
	mp[1] = "hi"
	mp[3] = "hey"
	mp[2] = "hello"
	mp[9] = "wow"
	mp[7] = "hurray"

	n := len(mp)
	arr := make([]int, n)
	i := 0
	for k := range mp {
		arr[i] = k
		i++
	}
	for j := 0; j < len(arr); j++ {
		println(arr[j])
	}
}

type number struct {
}
