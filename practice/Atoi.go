package piscine

import "os"

func Atoi(s string) int {
	neg := 1
	n := 0
	if s[0] == '-' {
		neg = -1
		s = s[1:]
	}
	for _, i := range s {
		if i < '0' || i > '9' {
			os.Exit(0)
		}
		n = n*10 + int(i-'0')
	}
	return neg * n
}
