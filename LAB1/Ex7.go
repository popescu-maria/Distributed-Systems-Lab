package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Eroare la deschiderea fisierului:", err)
		return
	}

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		counts[word] += 1
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire:", err)
		file.Close()
		return
	}
	file.Close()

	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}