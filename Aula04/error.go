package main

import (
	"fmt"
	"errors"
	"os"
)

func somarValores(numeros ...int) (int, error) {
	var soma int

	for _, valor := range numeros {
		if valor <= 0 {
			err := errors.New("Existe um valor negativo nos numeros")
			return 0, err
		}

		soma += valor
	}

	return soma, nil
}

func main() {

	result, err := somarValores(1, 2, 3, -1)
	if err != nil {
		fmt.Println("Erro ao somar os Valores")
		fmt.Println("A funcao 'somarValores()':", err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}
