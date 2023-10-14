package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr, err := strconv.Atoi(os.Args[1])
// 		if err == nil {
// 			if IsPowerOfTwo(nbr) {
// 				for _, w := range "true" {
// 					z01.PrintRune(w)
// 				}
// 			} else {
// 				for _, l := range "false" {
// 					z01.PrintRune(l)
// 				}
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func IsPowerOfTwo(n int) bool {
// 	return n != 0 && (n&(n-1) == 0)
// }
func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	if len(os.Args) == 2 {
		if Conv(num) {
			for _, s := range "true" {
				z01.PrintRune(s)
			}
		} else {
			for _, f := range "false" {
				z01.PrintRune(f)
			}
		}
	}
	z01.PrintRune('\n')
}

func Conv(n int) bool {
	if n%2 == 0 && n != 0 {
		return true
	} else {
		return false
	}
}
