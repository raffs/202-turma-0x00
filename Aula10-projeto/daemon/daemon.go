// arquivo: daemon/daemon
// para testar: go mod tidy (pega as lib; remove que nao esta sendo utilizado
//              go build (dentro do pacote para testar a compilacao)
package daemon

import (
	"errors"
	"context"
	"fmt"

	"github.com/4linux/4discovery/store"
	docker "github.com/docker/docker/client"
)

type Daemon struct {
	HTTPEndpoint string

	dockerClient *docker.Client
	dockerCtx    context.Context

	store store.ServiceStore
}

// Retorna uma nova instancia do Daemon{}
//   ex: daemon := daemon.New("localhost:8080")
func New(endpoint string) (d *Daemon, err error) {
	d = &Daemon{
		HTTPEndpoint: endpoint,
	}

	err = d.initDocker()
	if err != nil {
		fmt.Println("Erro ao iniciaalizar o docker")
		return nil, err
	}

	if ok := d.checkDocker(); !ok {
		fmt.Println("Nao foi possivel ao conextar o docker")
		return nil, errors.New("Nao foi possivel concetar ao docker")
	}

	d.store = store.New()
	return d, nil
}

func (d *Daemon) Start() error {
	comm := make(chan int)

	go d.servidorHTTP(comm)
	go d.dockerListener(comm)

	codigo := <-comm
	if codigo != 0 {
		return errors.New("Algum componente retornou um erro")
	}

	fmt.Println("Tchau!")
	return nil
}
