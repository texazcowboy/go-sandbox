package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			if f, err := os.Open(arg); err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
				f.Close()
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(stdin *os.File, counts map[string]int) {
	input := bufio.NewScanner(stdin)
	for input.Scan() {
		text := input.Text()
		if _, exists := counts[text]; exists {
			fmt.Println(stdin.Name())
		}
		counts[text]++
	}
}
