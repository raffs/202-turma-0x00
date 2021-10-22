package main

import (
	"fmt"
	"time"
)

func func4() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World 4")
}

func main() {
	go fmt.Println("Hello World 1")

	// go func() { ..... }()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello World 2")
	}() // realizando uma chamada da funcao.

	func3 := func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello World 3")
	}
	go func3()

	go func4()

	time.Sleep(1 * time.Second)
	fmt.Println("Hello World 5")
}
