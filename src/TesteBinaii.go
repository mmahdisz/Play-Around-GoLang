package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanln(&n)

	var trg string
	fmt.Scanln(&trg)

	var input string
	fmt.Scanln(&input)

	var r int
	for i := 0; i < n; i++ {
		if trg[i] != input[i] {
			r++
		}
	}
	fmt.Println(r)

}
