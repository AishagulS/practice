package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	len := len(os.Args)
// 	if len < 2 {
// 		z01.PrintRune('0')
// 	} else if len > 9 {
// 		num1 := len / 10
// 		num2 := len % 10
// 		z01.PrintRune(rune(num1 + '0'))
// 		z01.PrintRune(rune(num2 + '0'))
// 	} else {
// 		z01.PrintRune(rune(len + 48))
// 	}
// 	z01.PrintRune('\n')
// }
func main() {
	if len(os.Args[1:]) > 9 {
		z01.PrintRune(rune(len(os.Args[1:])/10 + 48))
		z01.PrintRune(rune(len(os.Args[1:])%10 + 48))
	} else {
		z01.PrintRune(rune(len(os.Args[1:]) + 48))
	}
	z01.PrintRune('\n')
}
