package main

import (
	"fmt"
	"sync"
)

func receiver(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	msg := <-ch
	fmt.Println("Mesaj primit in goroutina:", msg)
}

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	go receiver(ch, &wg)
	wg.Add(1)

	ch <- "salut"

	wg.Wait()
}