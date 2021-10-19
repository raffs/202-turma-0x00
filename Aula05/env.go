package main

import (
	"fmt"
	"os"
)

func main() {
	debugEnabled := os.Getenv("DEBUG_ENABLED")
	if debugEnabled == "" {
		debugEnabled = "true"
	}

	fmt.Println(debugEnabled)
}
