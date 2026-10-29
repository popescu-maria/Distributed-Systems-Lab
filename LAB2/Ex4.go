package main
import "fmt"

func gorutina(done chan bool) {
	fmt.Println("Goroutina...")
	done <- true
}

func main() {
	done := make(chan bool)

	go gorutina(done)

	<-done
	fmt.Println("Main: goroutina s a terminat")
}