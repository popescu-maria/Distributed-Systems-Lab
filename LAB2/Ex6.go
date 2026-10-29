package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(name string, finish int, done chan<- string) {
	position := 0
	for position < finish {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		position += 1 + rand.Intn(5)
		fmt.Printf("%s este la pozitia %d\n", name, position)
	}
	done <- name
}

func main() {
	const finishLine = 50

	ch1 := make(chan string)
	ch2 := make(chan string)

	fmt.Println("Observer: start cursa!")

	go robot("Robot 1", finishLine, ch1)
	go robot("Robot 2", finishLine, ch2)

	select {
	case winner := <-ch1:
		fmt.Printf("Observer: %s a ajuns primul la finish!\n", winner)
	case winner := <-ch2:
		fmt.Printf("Observer: %s a ajuns primul la finish!\n", winner)
	}
}