package main

import "github.com/01-edu/z01"

func main() {
	for _, i := range "Hello, world!" {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
