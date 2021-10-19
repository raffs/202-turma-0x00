package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func extrairEndereco(linha string) string {
	colunas := strings.Split(linha, " ")
	if len(colunas) < 1 {
		return ""
	}
	return strings.TrimSpace(colunas[0])
}

var usage string = `Uso %s [OPCOES]
Computar a quantidade de requisicoes por IP, dado um arquivo

OPCOES
   arquivo - Caminho de um arquivo de log (format apache ou nginx)
`

func main() {
	//linhas := []string{
	//	"8.8.8.8 [01010101] GET / ...",
	//	"8.8.8.4 [01010101] GET / ...",
	//	"8.8.8.8 [01010101] GET / ...",
	//	"8.8.8.8 [01010101] GET / ...",
	//	"8.8.8.4 [01010101] GET / ...",
	//	"8.8.8.8 [01010101] GET / ...",
	//}

	if len(os.Args) != 2 {
		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	caminhoArquivo := os.Args[1]
	arquivo, err := os.Open(caminhoArquivo)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err.Error())
		os.Exit(1)
	}
	defer arquivo.Close()

	//var acessos map[string]int
	// acessos = make(map[string]int)
	acessos := make(map[string]int)

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		endereco := extrairEndereco(linha)

		if _, existe := acessos[endereco]; existe {
			acessos[endereco] += 1
		} else {
			acessos[endereco] = 1
		}
	}

	for ip, req := range acessos {
		fmt.Println(ip, "=>", req)
	}
}
