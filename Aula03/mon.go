package main

import (
	"fmt"

	"github.com/raffs/go-labs/sistema"
)

func main() {
	usoCpu := sistema.CpuUso()

	// Existe um bug que quando o numero e 60, ele retorna a ATENCAO. Porem,
	// ele deveria retornar NORMAL :). 
	if usoCpu <= 60 {
		fmt.Println("NORMAL: A CPU esta com a utilizacao normal")
	} else if usoCpu < 85 {
		fmt.Println("ATENCAO: A CPU esta com a utilizacao alta")
	} else {
		fmt.Println("PERIGO: A CPU esta com a utilizaao execessiva")
	}

	// Ate 60 => normal
	// De 61 ate 85 => atencao
	// Maior que 85 => perigo
	usoMem := sistema.MemUso()

	if usoMem <= 60 {
		fmt.Println("NORMAL: Uso de memoria esta ok", usoMem)
	} else if usoMem < 80 {
		fmt.Println("ATENCAO: Uso de memoria esta alta", usoMem)
	} else {
		fmt.Println("PERIGO: Uso de memoria esta excessiva", usoMem)
	}
}
