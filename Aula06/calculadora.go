package main

import (
	"fmt"
	"strconv"
	"os"

	"4linux.com.br/4calc/mlib"
	"4linux.com.br/4calc/mlib/extra"
)

func main() {

	// <programa> add N N
	if len(os.Args) < 3 {
		fmt.Printf("uso: %s <command> <n2> <n2>\n", os.Args[0])
		os.Exit(2)
	}

	numero1, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("O valor '%s' nao pode ser convertido\n", os.Args[2])
		os.Exit(1)
	}

	numero2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("O valor '%s' nao pode ser convertido\n", os.Args[3])
		os.Exit(1)
	}

	operacao := os.Args[1]
	switch operacao {
	case "add":
		resultado := lib.Adicionar(numero1, numero2)
		fmt.Printf("%d + %d = %d\n", numero1, numero2, resultado)

		resultado1 := mlibv2.AdicionarV2(numero1, numero2)
		fmt.Printf("%d + %d = %d\n", numero1, numero2, resultado1)

	case "sub":
		resultado := lib.Substrair(numero1, numero2)
		fmt.Printf("%d - %d = %d\n", numero1, numero2, resultado)

	case "fact":
		fact := extra.Fatorial(numero1)
		fmt.Printf("%d! = %d\n", numero1, fact)

	default:
		fmt.Printf("O comando nao e suportado\n")
		os.Exit(1)
	}
}
