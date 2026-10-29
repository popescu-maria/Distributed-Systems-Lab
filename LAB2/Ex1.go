package main
import "fmt"

func mesaj(ch chan string) {
	ch <- "salut"
}

func main() {
	ch := make (chan string)

	go mesaj(ch)
	msj := <- ch
	fmt.Println("Mesaj primit in main: ", msj)

}