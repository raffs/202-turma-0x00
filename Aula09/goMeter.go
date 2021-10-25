package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
)

type Config struct {
	clientes  uint
	reqs      uint
}

type Sumario struct {
	reqSucesso uint
	reqErro    uint
	reqTotal   uint
}

func estressar(cfg *Config, site string) {
	// channel com buffer de `cfg.clientes` para armazenar este valor.
	chTotal := make(chan Sumario, cfg.clientes)

	for i := 0; i < int(cfg.clientes); i++ {
		go func() {
			// para cada cliente, faca:
			var sumParcial Sumario

			// cada cliente ganha sua goroutine ....
			for j := 0; j < int(cfg.reqs); j++ {
				// para cada requisicao, faca:

				resp, err := http.Get(site)
				if err != nil {
					sumParcial.reqErro += 1
				} else {
					if resp.StatusCode >= 200 && resp.StatusCode < 300 {
						sumParcial.reqSucesso += 1
					} else {
						sumParcial.reqErro += 1
					}
				}

				sumParcial.reqTotal += 1
			}

			chTotal <- sumParcial
			
		}() // chamando a funcao anonima().
	}

	var sumTotal Sumario

	// recebe um valor do Sumario (partial)
	for i := 0; i < int(cfg.clientes); i++ {
		sumParcial := <- chTotal

		sumTotal.reqSucesso += sumParcial.reqSucesso
		sumTotal.reqErro += sumParcial.reqErro
		sumTotal.reqTotal += sumParcial.reqTotal
	}

	fmt.Println("Sumario:", site)
	fmt.Println("QTDE De requisicoes com sucesso:", sumTotal.reqSucesso)
	fmt.Println("QTDE De requisicoes com erro:", sumTotal.reqErro)
	fmt.Println("QTDE De requisicoes no total:", sumTotal.reqTotal)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Uso: %s <site> <clientes> <qtde de requisicoes>\n")
		os.Exit(1)
	}

	site := os.Args[1]

	clientes, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Erro ao conventer:", os.Args[2])
		os.Exit(1)
	}

	reqs, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Erro ao conventer:", os.Args[3])
		os.Exit(1)
	}

	cfg := &Config{
		clientes: uint(clientes),
		reqs: uint(reqs),
	}

	estressar(cfg, site)
}
