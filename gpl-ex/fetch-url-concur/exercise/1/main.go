package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", secs)
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	if file, err := os.Create("output.html"); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		nbytes, err := io.Copy(file, resp.Body)
		resp.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	}
}
