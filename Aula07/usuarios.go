package main

import (
	"fmt"
)

type statusUsuario struct {
	habilitado  bool
	ultimoLogin string
	processos   uint
}

type usuario struct {
	nome, email string
	idade uint
	statusUsuario
}

func main() {
	var usuarioLogado usuario = usuario{
		nome: "Alice",
		email: "alice@4linux.com.br",
		idade: 25,
		statusUsuario: statusUsuario{
			habilitado: true,
			ultimoLogin: "21/10/2021 11:40",
			processos: 0,
		},
	}

	fmt.Println("Usuario: ", usuarioLogado)
}
