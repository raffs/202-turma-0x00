package main

import (
	"fmt"
	"strings"
)

//func (n Nome) primeiroNome() Nome {
//	nomes := strings.Split(string(n), "")
//	if len(nomes) < 0 {
//		return ""
//	}
//	return Nome(nomes[0])
//}

//func (n *Nome) caixaAlta() {
//	*n = Nome(strings.ToUpper(string(*n)))
//	fmt.Println(n)
//}
//

type Nome struct {
	primeiroNome, segundoNome string
}

func (n *Nome) caixaAlta() {
	n.primeiroNome = strings.ToUpper(n.primeiroNome)
	n.segundoNome  = strings.ToUpper(n.segundoNome)
}

func main() {
	usuario := Nome{
		primeiroNome: "Alice",
		segundoNome: "Silva",
	}

	fmt.Println("Nome do Usuario:", usuario)
	//fmt.Println("Primeiro Nome:", usuario.primeiroNome())

	usuario.caixaAlta()
	fmt.Println("Usuario:", usuario)
}
