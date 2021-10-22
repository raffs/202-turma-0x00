package main

import (
	"fmt"
	"os"

	"github.com/4linux/4monitor/daemon"
	"github.com/4linux/4monitor/api"
	"github.com/4linux/4monitor/plugins"
)

func usage() {
	fmt.Println(`Usage: 4monitor <command>

Comandos:

     versao    - Mostra a versao do daemon
     start     - Inicializar o daemon
`)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	command := os.Args[1]
	daemon := daemon.New("v0.0.1-alpha")

	// registrando todos os plugins que serao
	// codificados no pacote ./plugins
	daemon.Plugins = []api.Plugin{
		&plugins.Cpu{So: "linux"},
		&plugins.Cpu{So: "freebsd"},
	}

	switch command {
	case "versao":
		daemon.Versao()
	case "start":
		daemon.Start()
	default:
		fmt.Printf("O comando '%s' nao e suportado", command)
	}
}
