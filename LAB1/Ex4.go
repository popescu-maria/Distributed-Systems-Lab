package main

import (
	"fmt"
)

func main() {
	slice := []float32{10.5, 20.1, 30.8, 40.2, 50.9}
	fmt.Println("Slice-ul initial:", slice)

	var index int
	fmt.Print("Introduceti indexul elementului de eliminat: ")
	_, err := fmt.Scan(&index)

	if err != nil || index < 0 || index >= len(slice) {
		fmt.Println("Index invalid!")
		return
	}

	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println("Slice-ul dupa eliminare:", slice)
}