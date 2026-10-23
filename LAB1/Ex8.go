package main

import "fmt"

func main() {
	angajati := map[string]int{
		"Ana Popescu":       1,
		"Mihai Ionescu":     2,
		"Bogdan Radu":       1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":   2,
		"Ioana Enache":      1,
	}

	departamente := make(map[int][]string)

	for nume, idDep := range angajati {
		departamente[idDep] = append(departamente[idDep], nume)
	}

	for idDep, listaAngajati := range departamente {
		fmt.Printf("Departamentul %d: %v\n", idDep, listaAngajati)
	}
}