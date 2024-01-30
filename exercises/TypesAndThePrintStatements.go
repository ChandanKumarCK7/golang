package main

import (
	"fmt"
)

func main() {
	var student1 = "alpha"
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Print(student1, "\n")
	fmt.Print(a, b, c, d)
	fmt.Printf("%T\n", student1)
	fmt.Printf("%T\n", a)

	var x = 6
	x = int("6".toInt())
	fmt.Printf("%T\n", x)
}
