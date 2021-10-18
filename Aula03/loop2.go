package main

import (
	"fmt"
	"os"
)

func main() {

	// for i := 0; i < len(os.Args); i++ {
	// for _, parametro := range os.Args {
	for indice, parametro := range os.Args { 
		// fmt.Printf("Argumento[%d] = %s\n", i, os.Args[i])
		// fmt.Printf("Argumento[%d] = %s\n", indice, parametro)
		fmt.Printf("Argumento = %s\n", parametro)
	}
}
