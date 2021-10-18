package main

import (
	"fmt"
)

var versao string = "0.0.1"
var commit string = "f3drasd"

// Precisa: [INFO] texto ....
// Imprime uma linha nova \n
func info(formato string, opts ...interface{}) {
	msg := fmt.Sprintf(formato, opts...)
	fmt.Println("[INFO]", msg)
}

func critical(formato string, opts ...interface{}) {
	msg := fmt.Sprintf(formato, opts...)
	fmt.Println("[CRITICAL]", msg)

	panic("error critico no sistema")
}

func main() {
	info("Inicializando o servidor de paginas web (%s) commit (%s)", versao, commit)

	critical("Alguma coisa de errada aconteceu")
}
