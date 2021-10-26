package main

import (
	"fmt"
	"os"

	"github.com/4linux/4discovery/daemon"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Uso: %s <endereco de escuta>\n", os.Args[0])
		os.Exit(1)
	}
	
	url := os.Args[1]
	daemon, err := daemon.New(url)
	if err != nil {
		fmt.Println("Erro ao inicializar o daemon:", err.Error())
		os.Exit(1)
	}

	err = daemon.Start()
	if err != nil {
		fmt.Println("Erro durante o start() do daemon")
		os.Exit(1)
	}
}
