package main

import "fmt"

func main() {
	intArray := [5]int{1, 2, 3, 4, 5}
	intSlice := []int{}

	for i := range intArray {
		if intArray[i]%2 == 0 {
			intSlice = append(intSlice, intArray[i])
		}
	}

	fmt.Println(intSlice)

}
