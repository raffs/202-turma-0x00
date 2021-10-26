package main

import (
	"fmt"
)

func main() {
	slice1 := [2]string{"1", "2"}
	slice2 := [2]string{"1", "2"}
	slice3 := [2]string{"1", "3"}

	fmt.Println(slice1 == slice2)
	fmt.Println(slice1 == slice3)
}
