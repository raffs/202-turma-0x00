package main

import (
	"fmt"
)

var versao string = "ALPHA"

func alterarVersao() {
	versao := "0.0.1"
	fmt.Println("alterarVersao: ", versao)
}

func main() {
	alterarVersao()
	fmt.Println(versao)
}
