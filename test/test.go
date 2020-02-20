package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var a [][]uint8
	for I := 0; I < dy; I++ {
		var b []uint8
		for j := 0; j < dx; j++ {
			b = append(b, uint8(I)*uint8(j))
		}
		a = append(a, b)
	}
	return a
}

func main() {
	pic.Show(Pic)
}
