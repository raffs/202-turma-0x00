package main

import (
	"fmt"
)

func main() {
	// chave => do tipo string
	// valor => do tipo string
	usuarios := map[string]string{
		"Alice": "alice@4linux.com.br",
		"Bruno": "bruno@4linux.com.br",
		"Caio": "caio@4linux.com.br",
	}

	email, existe := usuarios["Alice"]
	if existe {
		fmt.Println("O usuario Alice existe no sistema", email)
	}

	if _, ok := usuarios["Danilo"] ; !ok {
		fmt.Println("O usuario Daniono nao existe no sistema")
	}

	
	fmt.Printf("MAPS: %v\n", usuarios)
}
