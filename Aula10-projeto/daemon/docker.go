package daemon

import (
	"context"
	"fmt"

	"github.com/4linux/4discovery/store"

	docker "github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func (d *Daemon) initDocker() (err error) {
	d.dockerCtx = context.Background()
	
	d.dockerClient, err = docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	return nil
}

func (d *Daemon) checkDocker() bool {
	info, err := d.dockerClient.Ping(d.dockerCtx)
	if err != nil {
		fmt.Println("ERRO: Ao tentar conectar docker:", err.Error())
		return false
	}

	fmt.Println("Conexao com docker realizada com sucesso:", info.APIVersion)
	return true
}

func (d *Daemon) addService(cid string) {
	container, err := d.dockerClient.ContainerInspect(d.dockerCtx, cid)
	if err != nil {
		fmt.Println("Nao foi possivel pegar info do container(containerInspect):", cid)
		return
	}

	name, existe := container.Config.Labels["4DISCOVERY_NAME"]
	if !existe {
		fmt.Printf("O container com o ID '%s' nao possui 4DISCOVERY_NAME, ignorando...")
		return
	}

	servico := store.Service{
		Name: name,
		ContainerId: cid,
		Image:       container.Config.Image,
		IPAddress:   container.NetworkSettings.IPAddress,
		Endpoint:    container.NetworkSettings.EndpointID,
		Hostname:    container.Config.Hostname,
		Labels:      container.Config.Labels,
	}

	fmt.Println("Adicionando servico:", servico)
	d.store.Add(name, servico)
}

func (d *Daemon) remService(cid string) {
	container, err := d.dockerClient.ContainerInspect(d.dockerCtx, cid)
	if err != nil {
		fmt.Println("Nao foi possivel pegar info do container(containerInspect):", cid)
		return
	}

	name, existe := container.Config.Labels["4DISCOVERY_NAME"]
	if !existe {
		fmt.Printf("O container com o ID '%s' nao possui 4DISCOVERY_NAME, ignorando...")
		return
	}

	fmt.Println("Removendo o servico:", name)
	d.store.Rem(name)
}

// Abre uma conexao com docker.events, escuta qualquer
// informacao. Baseando nela, ele adiciona ou remove
// o service.
func (d *Daemon) dockerListener(comm chan int) {
	dockerEventos, dockerErr := d.dockerClient.Events(d.dockerCtx, types.EventsOptions{})

	for {
		select {
		// tratamento do eventos que vem do Docker Daemon
		case event := <-dockerEventos:
			if event.Type == "container" {
				switch event.Status {
				case "create":
					d.addService(event.ID)

				case "die":
					d.remService(event.ID)
				}
			}

		// tratamento de errors do dockerEvents()
		case e := <-dockerErr:
			fmt.Println("Erro recebido pela docker daemon:", e.Error())
			comm <- 1
		}
	}

	comm <- 0
}
