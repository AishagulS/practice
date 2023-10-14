package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

// func FoldInt(f func(int, int) int, a []int, n int) {
// 	for i := 0; i < len(a); i++ {
// 		n = f(n, a[i])
// 	}
// 	for _, j := range Itoa(n) {
// 		z01.PrintRune(j)
// 	}
// 	z01.PrintRune('\n')
// }

// func Itoa(n int) string {
// 	neg := ""
// 	if n < 0 {
// 		neg = "-"
// 		n = -n
// 	}
// 	if n == 0 {
// 		z01.PrintRune('0')
// 	}
// 	str := ""
// 	for n > 0 {
// 		str = string(rune(n%10+48)) + str
// 		n = n / 10
// 	}
// 	return neg + str
// }

func Add(acc, cur int) int {
	return acc + cur
}

func Mul(acc, cur int) int {
	return acc * cur
}

func Sub(acc, cur int) int {
	return acc - cur
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	for _, v := range Itoa(n) {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Itoa(nb int) string {
	neg := ""
	var res string
	if nb < 0 {
		neg = "-"
		nb = -nb
	}
	if nb == 0 {
		z01.PrintRune('0')
	}
	if nb > 0 {
		res = string(rune(nb%10+48)) + res
		nb /= 10
	}
	return neg + res
}
