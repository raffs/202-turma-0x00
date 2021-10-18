package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: programa <numero>")
		os.Exit(1)
	}

	diaStr := os.Args[1]
	diaDaSemana, _ := strconv.Atoi(diaStr)

	switch diaDaSemana {
	case 0, 7:
		fmt.Println("Domingo")
	case 1:
		fmt.Println("Segunda")
	case 2:
		fmt.Println("Terca")
	case 3:
		fmt.Println("Quarta")
	case 4:
		fmt.Println("Quinta")
	case 5:
		fmt.Println("Sexta")
		fallthrough 
	case 6:
		fmt.Println("Sabado")

	default:
		fmt.Println("Dia nao reconhecido")
		os.Exit(1)
	}
}
