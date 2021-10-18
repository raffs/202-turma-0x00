package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	numeros := os.Args
	total := 0

	for indice, numeroStr := range numeros {
		numero, _ := strconv.Atoi(numeroStr)
		if indice == 0 {
			continue
		}
		total += numero
	}

	media := float64(total / len(numeros) - 1)

	fmt.Printf("Media: %.2f\n", media)
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Quantidade: %d\n", len(numeros))
}
