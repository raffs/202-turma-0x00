package main

import (
	"fmt"
	"time"
	"strconv"
	"os"

	"github.com/raffs/go-labs/sistema"
)

func main() {
	var totalCpu float64
	var totalMem float64

	if len(os.Args) < 3 {
		fmt.Println("Voce precisa especificar o <contador> <intervalo>")
		fmt.Printf("Uso: %s <contador> <intervalo>\n", os.Args[0])
		os.Exit(1)
	}

	contador, _ := strconv.Atoi(os.Args[1])
	intervalo, _ := strconv.Atoi(os.Args[2])

	for i := 0; i < contador; i++ {
		numeroCpu := sistema.CpuUso()
		totalCpu += numeroCpu

		numeroMem := sistema.MemUso()
		totalMem += numeroMem

		time.Sleep(time.Duration(intervalo) * time.Second)
	}

	mediaCpu := totalCpu / float64(contador)
	mediaMem := totalMem / float64(contador)

	fmt.Printf("Media: CPU=%.2f MEM=%.2f\n", mediaCpu, mediaMem)
	fmt.Printf("Total: CPU=%.2f MEM=%.2f\n", totalCpu, totalMem)
}
