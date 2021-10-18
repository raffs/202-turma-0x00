package main

import (
	"fmt"
)

func main() {
	var numero int = 60

	if numero > 100 {
		fmt.Println("O numero", numero, "e maior que 100")
	} else {
		fmt.Println("O numero", numero, "e menor que 100")
	}

	fmt.Println("Tchau!")
}
