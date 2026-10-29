package main
import "fmt"

func main() {
	ch := make(chan string, 2)

	go func() {
		ch <- "mesaj 1"
		ch <- "mesaj 2"
	}()

	msg1 := <-ch
	msg2 := <-ch

	fmt.Println("A fost primit:", msg1)
	fmt.Println("A fost primit:", msg2)
}