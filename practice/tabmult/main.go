package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num1, err := strconv.Atoi(args[0])
	if len(args) != 1 {
		return
	}
	if err == nil {
		if num1 < 0 {
			return
		}
		for num2 := 1; num2 < 10; num2++ {
			for _, w := range strconv.Itoa(num2) {
				z01.PrinRune(w)
			}
			z01.PrintRune(' ')
			z01.PrintRune('*')
			z01.PrintRune(' ')
			for _, r := range strconv.Itoa(num1) {
				z01.PrintRune(r)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			for _, k := range strconv.Itoa(num1 * num2) {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
