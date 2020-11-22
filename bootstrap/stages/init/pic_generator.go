package main

import (
	"golang.org/x/tour/pic"
	// "math"
)



func Pic(dx, dy int) [][]uint8 {
	var temp []uint8
	res := make([][]uint8, dy)
	for i, _ := range res {
		temp = make([]uint8, dx)
		res[i] = temp
		for j, _ := range temp {
			// res[i][j] = uint8((i+j)/2)
			// res[i][j] = uint8(math.Pow(float64(i),float64(j)))
			res[i][j] = uint8(i*j)
		}
	}
	return res
}

func main() {
	// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
	pic.Show(Pic) // https://github.com/golang/tour/blob/master/pic/pic.go // displays as base64
	// https://stackoverflow.com/questions/43212213/base64-string-decode-and-save-as-file
}
