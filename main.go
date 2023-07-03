package main

import (
	"fmt"
	"nmath/nmath"
)

func main() {
	a := nmath.Add(1, 4)
	fmt.Println("One plus two seems to be:", a)
	fmt.Println("One plus two seems to be:", a)
	a = nmath.Multiply(2, 2)
	fmt.Println("Two multiply two seems to be:", a)
}
