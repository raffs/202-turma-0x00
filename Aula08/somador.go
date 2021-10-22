package main

import (
	"fmt"

	"github.com/raffs/go-labs/numeros"
)

//                    / [1 2 4] --- goroutine \
// [ 1 2 3 4 5 6 7 ]  | [1 2 3] --- goroutine |+++ channel <--- somando
//                    \ [1 2 3] --- goroutine /

func soma(ch chan int, lista []int) {
	var resultado int

	for _, valor := range lista {
		resultado += valor
	}

	ch <- resultado
}

func main() {
	canal := make(chan int, 3)

	lista := numeros.GeradorLista(456789)
	subLista := numeros.DividirEm(lista, 3)

	// go soma(canal, subL[0])
	// go soma(canal, subL[1])
	// go soma(canal, subL[2])
	for _, subL := range subLista {
		go soma(canal, subL)
	}

	// var valor1 int
	valor1 := <-canal
	valor2 := <-canal
	valor3 := <-canal

	fmt.Println("Soma final:", valor1 + valor2 + valor3)
}
