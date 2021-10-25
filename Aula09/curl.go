package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Uso: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]

	resposta, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: Nao foi possivel ao conectar ao site: %s\n", err.Error())
		os.Exit(1)
	}

	defer resposta.Body.Close()
	conteudo, err := io.ReadAll(resposta.Body)
	if err != nil {
		fmt.Println("ERRO: erro ao ler o conteudo da requisicao:", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(conteudo))
}
