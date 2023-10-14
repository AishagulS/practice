// package main

// func main() {
// 	l := StrLen("r")
// 	fmt.Println(l)
// }

// func StrLen(s string) int {
// 	return len(s)
// }
package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	n := []rune(s)
	b := len(n)
	return b
}
