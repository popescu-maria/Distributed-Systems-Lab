package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	var prime bool = true

	if n <= 1 {
		prime = false
	}

	if n%2 == 0 {
		prime = false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println(n, "is a prime number.")
	} else {
		fmt.Println(n, "is not a prime number.")
	}
}
