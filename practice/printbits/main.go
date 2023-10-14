package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		res := ""
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		if num == 0 {
			res = "00000000"
		}
		for len(res) != 8 {
			a := num % 2
			res = string(rune(a+48)) + res
			num /= 2
		}
		for _, h := range res {
			z01.PrintRune(h)
		}
	}
	z01.PrintRune('\n')
}
