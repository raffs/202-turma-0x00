package main

import (
	"fmt"
	"strings"
)

func main() {
	// var palavra string
	// fmt.Println("E igual = ", string(letra) == palavra)

	// var letra rune = 'l'
	letra := 'l'
	fmt.Printf("%d %c\n", letra, letra)

	palavra := strings.ToUpper(strings.ToLower("4LINUX"))
	fmt.Println(palavra)
}
