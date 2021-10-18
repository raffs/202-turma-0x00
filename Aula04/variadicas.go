package main

import (
	"fmt"
)

func somarNumeros(numeros ...int) int {
	var resultado int

	for _, valor := range numeros {
		resultado += valor
	}

	return resultado
}

func main() {
	fmt.Println(somarNumeros(1, 2, 3, 4, 5))
	fmt.Println(somarNumeros(10, 20))
	fmt.Println(somarNumeros(1, 1, 1, 1, 1, 1, 1, 1, 1))
}
