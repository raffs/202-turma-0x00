package main

import (
	"fmt"
	"strconv"
	"os"
)

func fatorial(n int) (resultado int) {
	if n <= 1 {
		return 1
	}
	return fatorial(n - 1) * n
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Uso: %s <numero>\n", os.Args[0])
		os.Exit(1)
	}

	numero, err := strconv.Atoi(os.Args[1])
	// err = .... or err = nil
	if err != nil {
		fmt.Printf("O valor '%s' nao e numrero\n", os.Args[1])
		os.Exit(1)
	}
	
	resultado := fatorial(numero)
	fmt.Printf("%d! = %d\n", numero, resultado)
}
