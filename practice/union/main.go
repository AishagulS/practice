package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	args := os.Args
// 	if len(args) != 3 {
// 		fmt.Println()
// 		return
// 	}
// 	res := ""
// 	for _, i := range args[1] + args[2] {
// 		if !Do(res, i) {
// 			res += string(i)
// 		}
// 	}
// 	fmt.Println(res)
// }

// func Do(s string, r rune) bool {
// 	for _, c := range s {
// 		if c == r {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	a := os.Args[1:]
// 	if len(os.Args) != 3 {
// 		fmt.Println()
// 		return
// 	} else {
// 		res := ""
// 		for _, i := range a[0] + a[1] {
// 			if !Double(res, i) {
// 				res += string(i)
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

// func Double(s string, i rune) bool {
// 	for _, f := range s {
// 		if f == i {
// 			return true
// 		}
// 	}
// 	return false
// }
func main() {
	res := ""
	if len(os.Args) == 3 {
		a := os.Args[1]
		b := os.Args[2]
		for _, i := range a + b {
			if Double(res, i) {
				res += string(i)
			}
		}
		for _, d := range res {
			z01.PrintRune(d)
		}
	}
	z01.PrintRune('\n')
}

func Double(s string, r rune) bool {
	for _, j := range s {
		if j == r {
			return false
		}
	}
	return true
}
