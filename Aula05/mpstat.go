package main

import (
	"fmt"
	"strconv"
	"time"
	"os"

	"github.com/raffs/go-labs/sistema"
)

func processarArgs(params []string) (intervalo, contador uint) {
	intervalo = 1 * uint(time.Second)
	contador = 10

	switch len(params) {
	case 2:
		i, err := strconv.Atoi(params[1])
		if err != nil {
			fmt.Printf("O valor '%s' nao pode ser convertido\n", params[1])
			os.Exit(1)
		}
		intervalo = uint(i) * uint(time.Second) // vamos sobrescrever o valor padrao de `intervalo`

		fallthrough
	case 1:
		c, err := strconv.Atoi(params[0])
		if err != nil {
			fmt.Printf("O valor '%s' nao pode ser convertido\n", params[0])
			os.Exit(1)
		}
		contador = uint(c) // vamos sobrescrever o valor de `contador`

	default:
		// faz nada, segue o jogo!
	}

	return intervalo, contador
}

const usageMsg = `uso %s system [counter [intervalo]]

system
    Define qual parte do sistema sera monitorando: mem, cpu

opcoes
    counter   - Quantas vezes a utilizacao de CPU sera monitoranda
    intervalo - Intervalo entre cada requisicao
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(usageMsg, os.Args[0])
		os.Exit(1)
	}

	system := os.Args[1]
	intervalo, contador := processarArgs(os.Args[2:])

	switch system {
	case "cpu":
		for i := 0; i < int(contador); i++ {
			fmt.Println(sistema.CpuUso())
			time.Sleep(time.Duration(intervalo))
		}
	case "mem":
		for i := 0; i < int(contador); i++ {
			fmt.Println(sistema.MemUso())
			time.Sleep(time.Duration(intervalo))
		}
	default:
		fmt.Printf("Sistema '%s' nao existe ou nao e suportado\n", system)
		fmt.Printf(usageMsg, os.Args[0])
	}

}
