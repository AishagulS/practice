package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			nb = FindPrevPrime(nb - 1)
// 		}
// 	}
// 	return nb
// }

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for nb > 1 {
		if Prev(nb) {
			break
		}
		nb--
	}
	return nb
}

func Prev(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
