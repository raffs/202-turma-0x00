package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("./paginas"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor de paginas estaticas")
		os.Exit(1)
	}
}
