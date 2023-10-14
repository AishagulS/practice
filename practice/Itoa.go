package piscine

import "fmt"

func Itoa(n int) string {
	res := ""
	neg := ""
	if n < 0 {
		neg = "-"
		n = -n
	}
	if n == 0 {
		fmt.Println(0)
	}
	if n > 0 {
		res = string(rune(n%10)+48) + res
		n = n / 10
	}
	return neg + res
}
