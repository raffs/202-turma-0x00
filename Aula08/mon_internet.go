package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func verificarConn(site string) {
	fmt.Println("Verificando a conexao da internet com", site)

	_, err := net.Dial("tcp", site)
	if err != nil {
		fmt.Println("Falhou:", err.Error())
	} else {
		fmt.Println("Ok")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Uso: %s <site>\n", os.Args[0])
		os.Exit(1)
	}
	site := os.Args[1]

	// ticker <<<<< Manda um sinal a cada N tempo
	// signal <<<<< Receber sinal SO, finaliza programa com estilo :)

	ticker := time.NewTicker(1 * time.Second)

	chSignalSO := make(chan os.Signal)
	signal.Notify(chSignalSO, os.Interrupt)

	for {
		select {
		case sig := <-chSignalSO:
			fmt.Println("Recebi um sinal do SO, vou finalizar a minha exec aqui")
			fmt.Println("Signal:", sig)
			os.Exit(1)

		case <-ticker.C:
			verificarConn(site)
		}
	}
}
