package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var m int
	fmt.Scanln(&m)

	var r int64
	for i := -10; i <= m; i++ {
		for j := 1; j <= n; j++ {
			x := float64(i + j)
			pow := math.Pow(x, 3)
			f := math.Pow(float64(j), 2)
			u := pow / f
			r = r + int64(u)
		}
	}
	fmt.Println(r)

}
