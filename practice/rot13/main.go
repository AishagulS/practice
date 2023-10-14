package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'A' && i <= 'Z' {
				if i >= 'A'+13 {
					i -= 13
				} else {
					i += 13
				}
			} else if i >= 'a' && i <= 'z' {
				if i >= 'a'+13 {
					i -= 13
				} else {
					i += 13
				}
			}
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
