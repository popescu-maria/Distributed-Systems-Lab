package main

import "fmt"

func main() {
	m := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	names := make([]string, 0, len(m))
	heights := make([]int, 0, len(m))

	for name, height := range m {
		names = append(names, name)
		heights = append(heights, height)
	}

	n := len(heights)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}

	for i := 0; i < len(names); i++ {
		fmt.Printf("%s: %d\n", names[i], heights[i])
	}
}