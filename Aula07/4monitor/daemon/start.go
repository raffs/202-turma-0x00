// arquivo: daemon/start.go
package daemon

import (
	"fmt"
	"time"
	"sync"

	"github.com/4linux/4monitor/api"
)

func (d *Daemon) initPlugins() {
	for _, plugin := range d.Plugins {
		plugin.Iniciar()
	}
}

func executarColeta(p api.Plugin, wg *sync.WaitGroup) {
	err := p.Coletar()
	if err != nil {
		fmt.Println("ERRO ao coletar o plugin:", err)
	}
	wg.Done()
}

func (d *Daemon) Start() {
	var wg sync.WaitGroup

	fmt.Println("Inicilizando os plugins")
	d.initPlugins()

	for { // laco inifito
		wg.Add(len(d.Plugins)) // N plugins registrado no daemon.

		for _, plugin := range d.Plugins {
			// go func(p api.Plugin) {
			//	err := p.Coletar()
			//	if err != nil {
			//		fmt.Println("ERRO ao coletar o plugin:", err)
			//	}
			//	wg.Done()
			//}(plugin) // realizar a chamda da funcao
			go executarColeta(plugin, &wg)
		}

		wg.Wait()
		time.Sleep(1 * time.Second)
	}
}
