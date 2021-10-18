package main

import (
	"fmt"
)

func main() {
	numeros1 := []{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(quebrarNoMeio(numeros1))

	numeros2 := []{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(quebrarNoMeio(numeros2))
}
