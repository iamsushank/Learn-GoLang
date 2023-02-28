package main

func Pic(dx, dy int) [][]uint8 {
	var a [][]uint8
	for i := 0; i < dy; i++ {
		var b []uint8
		for j := 0; j < dx; j++ {
			b = append(b, uint8(i*j))
		}
		a = append(a, b)
	}
	return a
}

//
//func main() {
//	pic.Show(Pic)
//}
