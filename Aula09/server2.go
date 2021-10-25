package main

import (
	"fmt"
	"os"
	"net/http"
)

func getHtml() string {
	return `
<html>
  <head>
    <title> Hello World </title>
  </head>
  <body background="#F0F0F0">
     <h1> Hello World </h1>
  </body>
</html>
`
}

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		conteudoHtml := getHtml()
		fmt.Fprintf(w, conteudoHtml)
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
