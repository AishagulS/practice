package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) == 4 {
// 		s := os.Args[1]
// 		s1 := os.Args[2]
// 		s2 := os.Args[3]
// 		for _, i := range s {
// 			if i == rune(s1[0]) {
// 				z01.PrintRune(rune(s2[0]))
// 			} else {
// 				z01.PrintRune(rune(i))
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

func main() {
	args := os.Args[1:]
	a := args[0]
	b := args[1]
	c := args[2]
	if len(args) != 3 {
		return
	}
	for _, i := range a {
		if i == rune(b[0]) {
			z01.PrintRune(rune(c[0]))
		} else {
			z01.PrintRune(rune(i))
		}
	}
	z01.PrintRune('\n')
}
