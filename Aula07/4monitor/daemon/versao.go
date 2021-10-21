package daemon

import (
	"fmt"
)

func (d *Daemon) Versao() {
	fmt.Printf("4monitor versao %s \n", d.versao)
}
