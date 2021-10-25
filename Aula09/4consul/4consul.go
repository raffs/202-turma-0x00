package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Servico struct {
	Name     string
	Endpoint string
	Labels   map[string]string // type Labels map[string]string; Labels Labels
}

type ServicoStore struct {
	servicos []Servico

	sync.Mutex
}

var servicoStore ServicoStore

func (s *ServicoStore) Add(novoServico Servico) {
	s.Lock()
	s.servicos = append(s.servicos, novoServico)
	s.Unlock()
}

func servicoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if len(servicoStore.servicos) <= 0 {
			fmt.Fprintf(w, "Lista vazia")
			return
		}

		servicos, err := json.Marshal(servicoStore.servicos)
		if err != nil {
			fmt.Println("Nao foi possivel converter a estrutura servicos para json")
			fmt.Fprintf(w, "500 Error")
			return
		}

		fmt.Fprintf(w, string(servicos))

	case "POST":
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Nao foi possivel receber os conteudo da requisicao")
			fmt.Fprintf(w, "500 Error")
			return
		}
		defer r.Body.Close()

		var srv Servico
		if err := json.Unmarshal(payload, &srv); err != nil {
			fmt.Fprintf(w, "Nao foi possivel converter o payload em json")
			return
		}

		servicoStore.Add(srv)
		fmt.Fprintf(w, "OK!")
	}
}

func main() {
	http.HandleFunc("/servicos", servicoHandler)

	fmt.Println("Iniciando o servidor de catalog de servicos")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nao foi possivel iniciar o servidor http:", err.Error())
		os.Exit(1)
	}
}
