package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Command: %v", os.Args[0])
	fmt.Println()
	fmt.Printf("Args: %v", strings.Join(os.Args[1:], " "))
}
