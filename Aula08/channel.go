package main

import (
	"fmt"
	"time"
)

func main() {
	//var canal chan int
	canal := make(chan int)

	// envia um valor 5 (int)
	go func() {
		time.Sleep(5 * time.Second)
		canal <- 5
	}()

	// ------------- PIPE ------------------------- /

	// recebimento deste valor, joga para variavel resultado
	resultado := <- canal
	fmt.Println(resultado)

	// sincronizando as goroutines, mas nao utilizando os dados. Apenas esperando!
	<-canal
}
