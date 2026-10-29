package main

import "fmt"

func send(ch1 chan<- string) {
	ch1 <- "Mesaj de la send"
}

func receive(ch1 <-chan string, ch2 chan<- string) {
	msg := <-ch1
	ch2 <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1)
	go receive(ch1, ch2)

	rezultat := <-ch2
	fmt.Println("Main a afisat:", rezultat)
}