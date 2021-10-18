package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"os"
)

func main() {
	// validacao para os argumentos
	programa := os.Args[0]
	if len(os.Args) < 2 {
		fmt.Println("Voce precisa especificar um numero como argumento")
		fmt.Printf("Uso: %s <numero>\n", programa)
		os.Exit(1) // finalize programa, e retorna codigo 1
	}

	numeroStr := os.Args[1]
	numeroInt, _ := strconv.Atoi(numeroStr)

	rand.Seed(time.Now().UnixNano())
	somador := rand.Intn(1000)

	numero := numeroInt + somador

	if numero % 2 == 0 {
		fmt.Printf("O numero %d e par!\n", numero)
	} else {
		fmt.Printf("O numero %d nao e par!\n", numero)
	}
}
