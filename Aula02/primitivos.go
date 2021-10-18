package main

import (
	"fmt"
)

const (
	numero2 int32 = 1235
)

func main() {
	var numero int32 = 0x7FFFFFFF 
	// int8, int16, int32, int64 (int) int32, int64
	//var positivo uint = 0x8FFFFFFF

	// uint8, uint16, uint32, uint64 
	var gravidade float32 = 9.8
	//var pi float64 = 3.1415

	//fmt.Println("Numero =", numero)
	//fmt.Println("Gravidadde =", gravidade)
	//fmt.Println("Pi", pi)
	fmt.Println("Positivo =", numero2)	
	fmt.Println("Verdadeiro =", float32(numero) > gravidade)
}
