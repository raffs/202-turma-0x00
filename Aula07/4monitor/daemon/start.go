// arquivo: daemon/start.go
package daemon

import (
	"fmt"
	"time"
)

func (d *Daemon) initPlugins() {
	for _, plugin := range d.Plugins {
		plugin.Iniciar()
	}
}

func (d *Daemon) Start() {
	fmt.Println("Inicilizando os plugins")
	d.initPlugins()

	for { // laco inifito
		for _, plugin := range d.Plugins {
			err := plugin.Coletar()
			if err != nil {
				fmt.Println("ERRO ao coletar o plugin:", err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
