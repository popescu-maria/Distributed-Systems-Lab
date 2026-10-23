package main
import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite", "Bla"}
	fmt.Println("Slice ul initial: ", slice)

	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n - 1 - i] = slice[n - 1 - i], slice[i]
	}

	fmt.Println("Slice inversat: ", slice)
}