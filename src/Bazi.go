package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(1000)
	//println("x", x)
	var n int

MainLoop:
	for i := 25; i > 0; i-- {
		println("chance", i)
		fmt.Scanln(&n)
		switch {
		case n < x:
			println("go up!")
		case n > x:
			println("go down!")
		default:
			break MainLoop
		}
	}

}
