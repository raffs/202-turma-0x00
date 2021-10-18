package main

import (
	"fmt"
)

func main() {
	media, soma := calcularMedia([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Printf("Media=%d | Soma=%d\n", media, soma)
}
