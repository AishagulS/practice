package main

import (
	"github.com/01-edu/z01"
)

// func main() {
// 	z01.PrintRune(FirstRune("Hello!"))
// 	z01.PrintRune(FirstRune("Salut!"))
// 	z01.PrintRune(FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// func FirstRune(s string) rune {
// 	return rune(s[0])
// }
func FirstRune(s string) rune {
	n := []rune(s)
	return n[0]
}

func main() {
	z01.PrintRune(FirstRune("假借"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("la!"))
	z01.PrintRune('\n')
}
