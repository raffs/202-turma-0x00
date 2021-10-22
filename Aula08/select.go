package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

//    canal (sincronizao) < dados > .......
//    canal tempo (timeout) < dados < (dormir N segundos)

func worker(chResultado chan int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int() % 25
	fmt.Println("Aguardando:", r)

	time.Sleep(time.Duration(r) * time.Second)
	chResultado <- r
}

func timeout(chTimeout chan bool, tempo uint) {
	time.Sleep(time.Duration(tempo) * time.Second)
	chTimeout <- true
}

func main() {
	chResultado := make(chan int)
	chTimeout := make(chan bool)

	go worker(chResultado)
	go timeout(chTimeout, 10)

	select {
	case res := <-chResultado:
		fmt.Println("Resultado da conta e:", res)
	case <-chTimeout:
		fmt.Println("Error: timeout() ao executar o worker()")
		os.Exit(1)
	}
}
