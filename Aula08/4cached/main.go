package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"net"
)

func processarConn(conn net.Conn) {
	defer conn.Close()

	dados, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Nao foi possivel ler os dados")
		return
	}

	// ADD, chave, valor
	// DEL, chave
	linha := strings.Split(strings.Trim(dados, "\n"), " ")
	if len(linha) < 1 {
		fmt.Printf("Erro ao processar o protocol")
		return
	}

	comando := linha[0]
	switch comando {
	case "ADD":
		chave := linha[1]
		valor := linha[2]
		fmt.Printf("[S] ADD %s => %s\n", chave, valor)
		gCacheD.Add(chave, valor)
		fmt.Fprint(conn, "OK\n")

	case "DEL":
		chave := linha[1]
		fmt.Printf("[S] DEL %s\n", chave)
		gCacheD.Del(chave)
		fmt.Fprint(conn, "OK\n")

	case "GET":
		chave := linha[1]
		fmt.Printf("[S] GET %s\n", chave)
		valor := gCacheD.Get(chave)
		fmt.Fprint(conn, valor)

	default:
		fmt.Printf("[S] Comando %s nao e valido\n", comando)
		fmt.Fprint(conn, "FAIL\n")
	}
}

func main() {
	gCacheD = CacheD{
		d: make(map[string]string),
	}

	servidor, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor %s\n",
			err.Error())
		os.Exit(1)
	}

	fmt.Println("[S] Servidor escutando conexoes")

	for {
		conn, err := servidor.Accept()
		if err != nil {
			fmt.Println("[S] Erro ao receber a conexao")
			continue
		}

		go processarConn(conn)
	}
	
}
