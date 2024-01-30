package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array a", a)

	var b = [...]int{6, 7, 8, 9, 14}
	fmt.Println("Array b before reassinging", a)
	b[0] = 1

	fmt.Println("oth element after reassigning", b[0])
	fmt.Println("o-4th elements", b[0:4])

	fmt.Println("length of the array", len(b))
}
