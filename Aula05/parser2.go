package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func parsearLinha(linha string, delim string) (string, string) {
	// " " [GET] [/]
	// ":" [Host] [Value]
	colunas := strings.Split(linha, delim)
	if len(colunas) != 2 {
		fmt.Println("WARNING:")
		return "", ""
	}

	return strings.TrimSpace(colunas[0]), strings.TrimSpace(colunas[1])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// linhas := make([]string, 0)

	scanner.Scan()
	linha0 := scanner.Text()
	metodo, caminho := parsearLinha(linha0, " ")

	fmt.Println("METODO HTTP:", metodo)
	fmt.Println("DOCUMENTO:", caminho)

	fmt.Println("------")
	for scanner.Scan() {
		linha := scanner.Text()
		param, valor := parsearLinha(linha, ":")
		fmt.Println(param, "=>", valor)
	}
}
