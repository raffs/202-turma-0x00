package main

import (
	"fmt"
	"sort"
	"strconv"
	"os"
)

func maiorEm(numeros []int) int {
	if len(numeros) <= 0 {
		return 0
	}

	maior := numeros[0]
	for _, valor := range numeros {
		if valor > maior {
			maior = valor
		}
	}

	return maior
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Uso: %s <numero1> [numero2] ....\n", os.Args[0])
		os.Exit(1)
	}

	var numeros []int
	for _, valorStr := range os.Args[1:] {
		num, err := strconv.Atoi(valorStr)
		if err != nil {
			fmt.Printf("WARING: O valor '%s' na pode ser convertido para numero\n", valorStr)
			continue
		}

		numeros = append(numeros, num)
	}

	// fmt.Println(numeros)
	sort.Ints(numeros)
	// fmt.Println(numeros)

	// 1. Computar o menor
	// 2. Computar a media
	// 3. Computar o mediano // [1 *2* 3] (se a qtde for impar, pegar o valor do meio
	//                          [1 *2 3* 4] (se a qtde for par, media dos dois valores do meio)

	maior := maiorEm(numeros)
	fmt.Println("Maior numero e", maior)
}
