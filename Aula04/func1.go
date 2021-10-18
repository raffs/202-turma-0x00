package main

import (
	"fmt"
)

func somar(a int, b int) int {
	resultado := a + b
	return resultado
}

func main() {
	resultado := somar(1, 10)
	fmt.Printf("1 + 10 = %d\n", resultado)
}
