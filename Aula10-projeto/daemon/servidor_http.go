package daemon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (d *Daemon) services(w http.ResponseWriter, r *http.Request) {
	services := d.store.List()

	payload, err := json.Marshal(services)
	if err != nil {
		fmt.Println("Nao foi possivel conveter o catalogo de servicos para json")
		fmt.Fprintf(w, "FALHA!")
		return
	}

	// [
	//    {nome: XXXX, containerId, XXXX},
	//    { ... }
	//    { ... }
	//]
	fmt.Fprintf(w, string(payload))
}

func (d *Daemon) status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func (d *Daemon) servidorHTTP(comm chan int) {
	http.HandleFunc("/status", d.status)
	http.HandleFunc("/services", d.services)

	fmt.Println("Servidor HTTP iniciando ...")
	err := http.ListenAndServe(d.HTTPEndpoint, nil)
	if err != nil {
		fmt.Println("[S] Erro ao iniciar o servidor http")
		comm <- 1
	}

	comm <- 0
}
