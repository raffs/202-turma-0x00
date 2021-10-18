package main

import (
	"fmt"

	"github.com/raffs/go-labs/fs"
)

func parsearLinha(linha string, delim rune) (pos int) {
	for indice, valor := range linha {
		if valor == delim {
			pos = indice + 1
			break
		}
	}

	fmt.Println("Processando linha: ", linha)
	fmt.Println("Espaco na posicao: ", pos)

	// return
	return pos
}

func main() {
	// linha0 := "GET /"
	// linha1 := "Host: 4linux.com.br"

	linhas, err := fs.LerArquivo("request.txt")
	if err != nil {
		fmt.Println("Erro ao tentar ler o arquivo: ", err.Error())
	}

	// 1. Identificar qual linha temos que processar de acordo o delimitador:
	//      Primeira linha: GET /
	//      Outras linhas: host: 4linux ...
	//
	// 2. Parsar o arquivo `request.txt` por linha de comando
	//
	// 3. (Desafio) Varios request HTTP, no mesmo payload
	//    GET /
	//    Host: 4linux
	//    ...
	//    <linha vazia>
	//    GET /
	//    Host: 4linux.com.br
	//    <linha vazia>
	for linhaNum, linhaTexto := range linhas {
		fmt.Printf("Processando linha #%d: %s\n", linhaNum, linhaTexto)
	}
}
