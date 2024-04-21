package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Pirate Bot")
	fmt.Println("---------------------")
	err := getOrSetKey()
	if err != nil {
		os.Exit(1)
	}
}
