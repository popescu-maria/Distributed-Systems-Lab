package main

import "fmt"

func main() {
	var n, s int
	fmt.Print("Introduceti numarul de elemente si suma s: ")
	fmt.Scan(&n, &s)

	arr := make([]int, n)
	fmt.Println("Introduceti elementele: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	m := make(map[int]int)

	for i, val := range arr {
		complement := s - val
		idx, ok := m[complement]
		if ok {
			fmt.Println("Rezultat: ", idx, i)
		}
		m[val] = i
	}
}