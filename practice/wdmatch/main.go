package main

import (
	"fmt"
	"os"
)

// func main() {
// 	Args := os.Args
// 	if len(Args) != 3 {
// 		return
// 	}
// 	a1 := os.Args[1]
// 	a2 := os.Args[2]
// 	var count int
// 	var s string
// 	for i := 0; i < len(a1); i++ {
// 		for j := count + 1; j < len(a2); j++ {
// 			if a1[i] == a2[j] {
// 				s += string(a1[i])
// 				j = len(a2) - 1
// 			}
// 			count++
// 		}
// 	}
// 	if s == a1 {
// 		for _, b := range a1 {
// 			z01.PrintRune(b)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args
// 	if len(args) == 3 {
// 		s1 := args[1]
// 		s2 := args[2]
// 		j := 0
// 		for i := 0; i < len(s2); i++ {
// 			if j >= len(s1) {
// 				break
// 			}
// 			if s2[i] == s1[j] {
// 				j++
// 			}
// 		}
// 		if j == len(s1) {
// 			fmt.Println(s1)
// 		}
// 	}
// }
// func main() {
// 	args := os.Args
// 	if len(args) == 3 {
// 		s1 := args[1]
// 		s2 := args[2]
// 		j := 0
// 		for i := 0; i < len(s2); i++ {
// 			if s2[i] == s1[j] {
// 				j++
// 			}
// 			if j >= len(s1) {
// 				break
// 			}

// 		}
// 		if j == len(s1) {
// 			fmt.Println(s1)
// 		}
// 	}
// }
func main() {
	a := os.Args[1]
	b := os.Args[2]
	j := 0
	if len(os.Args) == 3 {
		for i := 0; i < len(b); i++ {
			if b[i] == a[j] {
				j++
			}
			if j >= len(a) {
				break
			}
		}
		if j == len(a) {
			fmt.Println(a)
		}
	}
}
