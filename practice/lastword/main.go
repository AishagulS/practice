package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	a := os.Args[1:]
// 	var res string
// 	if len(a) != 1 {
// 		return
// 	}
// 	for i := len(a[0]) - 1; i >= 0; i-- {
// 		if a[0][i] != ' ' {
// 			res = string(a[0][i]) + res
// 		}
// 		if res != "" && a[0][i] == ' ' {
// 			break
// 		}
// 	}
// 	if res == "" {
// 		return
// 	}
// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }
func main() {
	if len(os.Args) == 2 {
		res := ""
		a := os.Args[1]
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] != ' ' {
				res = string(a[i]) + res
			}
			if a[i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
