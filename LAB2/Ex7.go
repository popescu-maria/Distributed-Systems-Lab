package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Point struct {
	x, y int
}

func robot(name string, steps int, out chan<- Point, in <-chan Point, done chan<- string) {
	received := make(chan struct{})

	go func() {
		for p := range in {
			fmt.Printf("  %s a primit pozitia (%d, %d)\n", name, p.x, p.y)
		}
		fmt.Printf("%s: celalalt robot s-a oprit\n", name)
		close(received)
	}()

	x, y := 0, 0
	for i := 0; i < steps; i++ {
		time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
		x += rand.Intn(3)
		y += rand.Intn(3)
		out <- Point{x, y}
	}
	fmt.Printf("%s si-a terminat traseul\n", name)
	close(out)

	<-received

	done <- name
}

func main() {
	ch12 := make(chan Point) 
	ch21 := make(chan Point)
	done := make(chan string)

	fmt.Println("Observer: pornesc robotii")
	go robot("Robot 1", 5, ch12, ch21, done)
	go robot("Robot 2", 8, ch21, ch12, done)

	for i := 0; i < 2; i++ {
		fmt.Printf("Observer: %s a raportat ca a terminat\n", <-done)
	}
	fmt.Println("Observer: toti robotii au terminat")
}