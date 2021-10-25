package main

import (
	"fmt"
	"os"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello World")
		fmt.Println("[S] Acessou!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo nao suportado")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", index)
	http.HandleFunc("/ola", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nao foi possivel criar um servidor um servider http:", err.Error())
		os.Exit(1)
	}
}
