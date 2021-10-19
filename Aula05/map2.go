package main

import (
	"fmt"
)

// type <nome> <tipo>
type Usuario map[string]string
type Usuarios map[string]Usuario

func pegarUsuario(login string, usuarios Usuarios) (u Usuario) {
	if _, ok := usuarios[login]; !ok {
		return nil
	}

	u = usuarios[login]
	return
}

func main() {
//	usuarios := map[string]Usuario{
//		"alice": Usuario{
//			"email": "alice@4linux.com.br",
//			"telefone": "123456789",
//		},
//		"bob": Usuario{
//			"email": "bob@4linux.com.br",
//			"telefone": "133456789",
//		},
//	}

	usuarios := Usuarios{
		"alice": Usuario{
			"email": "alice@4linux.com.br",
			"telefone": "123456789",
		},
		"bob": Usuario{
			"email": "bob@4linux.com.br",
			"telefone": "133456789",
		},
	}
	
	a := pegarUsuario("alice", usuarios)
	fmt.Println("Alice:", a)

	d := pegarUsuario("danilo", usuarios)
	if d == nil {
		fmt.Println("Danilo nao existe")
	} else {
		fmt.Println("Danilo:", d)
	}
}
