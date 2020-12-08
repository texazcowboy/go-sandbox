package main

import (
	"fmt"
	"os"
)

func main() {
	var finalStr, delimiter string
	for i, arg := range os.Args {
		finalStr += fmt.Sprintf("%vindex: %v value:%v", delimiter, i, arg)
		delimiter = " "
	}
	fmt.Println(finalStr)
}
