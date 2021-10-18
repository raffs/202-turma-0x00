package main

import (
	"fmt"
)

func main() {
	// ... passa o trabalho de contar os elementos para o compilador
	numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// soma, maiorValor := extrairSomaEMaior(numeros)
	fmt.Println("[1:]  |", numeros[1:])
	fmt.Println("[:7]  |", numeros[:7])
	fmt.Println("[2:6] |", numeros[2:6])

	fmt.Println(len(numeros), cap(numeros))
}
