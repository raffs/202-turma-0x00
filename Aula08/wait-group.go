package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func worker() {
	randomId := rand.Intn(5000)
	fmt.Printf("Worker[%d] executando o trabalho ...\n", randomId)

	randomT := rand.Intn(100) % 5 
	time.Sleep(time.Duration(randomT) * time.Second)
	fmt.Printf("Worker[%d] finalizou o seu trabalho ...\n", randomId)
}

// wg.Contador De Rotinas (10)

// 1. Rotina #1 (1)
// 2. Rotina #2 (2)
// 3. Rotina #3 (3)

// 1. Rotina #2 wg.Done() (3 - 1) = 2
// 2. Rotina #1 wg.Done() (2 - 1) = 1
// 3. Rotina #3 wg.Done() (1 - 1) = 0

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			worker()
			wg.Done()
		}() // chamar a funcao
	}

	wg.Wait() // espera as rontinas finalizarem
	fmt.Println("Tchau!")
}
