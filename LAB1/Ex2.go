package main

import "fmt"

func main() {
	array := []int{4, 8, 2, 2, 3, 5, 9, 2, 10, -2}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)
}
