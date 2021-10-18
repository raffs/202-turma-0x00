package main

import (
	"fmt"
	"time"

	"github.com/raffs/go-labs/agendador"
)

func main() {
	fmt.Println("Inicializando o agendador de tarefa")

	tarefas, _ := agendador.LerConfigAgendador("agendador.conf")

	for {
		agora := time.Now()
		segundoAtual := agora.Second()
		minutoAtual := agora.Minute()
		horaAtual := agora.Hour()

		for _, tarefa := range tarefas {
			if tarefa.Minuto == minutoAtual {
				fmt.Printf("[Minuto=%d] Executando comando %s\n", minutoAtual, tarefa.Command)
			}

			if tarefa.Segundo == segundoAtual {
				fmt.Printf("[Segundo=%d] Executando comando %s\n", segundoAtual, tarefa.Command)
			}

			if tarefa.Hora == horaAtual {
				fmt.Printf("[hora=%d] Executando comando %s\n", horaAtual, tarefa.Command)
			}
		}

		time.Sleep(200)
	}
}
