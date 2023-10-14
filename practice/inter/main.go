package main

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 3 {
// 		fmt.Println()
// 		return
// 	}
// 	a := args[0]
// 	b := args[1]
// 	res := ""
// 	for _, i := range a {
// 		for _, v := range b {
// 			if i == v && !Do(res, i) {
// 				res += string(i)
// 			}
// 		}
// 	}
// 	fmt.Println(res)
// }

// func Do(res string, r rune) bool {
// 	for _, w := range res {
// 		if r == w {
// 			return true
// 		}
// 	}
// 	return false
// }
// package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	a := os.Args[1:]
// 	if len(a) != 2 {
// 		fmt.Println()
// 		return
// 	} else {
// 		res := ""
// 		for d := 0; d < len(a[0]); d++ {
// 			for t := 0; t < len(a[1]); t++ {
// 				if a[0][d] == a[1][t] {
// 					if Checker(res, rune(a[0][d])) {
// 						res += string(a[0][d])
// 					}
// 				}
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

// func Checker(s string, r rune) bool {
// 	for _, k := range s {
// 		if k == r {
// 			return false
// 		}
// 	}
// 	return true
// }
func main() {
	if len(os.Args) == 3 {
		res := ""
		a := os.Args[1]
		b := os.Args[1]
		for _, i := range a {
			for _, f := range b {
				if i == f {
					if Double(res, i) && i != ' ' {
						res += string(i)
					}
				}
			}
		}
		for _, h := range res {
			z01.PrintRune(h)
		}
	}
	z01.PrintRune('\n')
}

func Double(s string, r rune) bool {
	for _, k := range s {
		if k == r {
			return false
		}
	}
	return true
}
