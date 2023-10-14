package main

import (
	"fmt"
	"os"
)

// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	a := Atoi(os.Args[1])
// 	b := Atoi(os.Args[3])
// 	if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
// 		return
// 	}
// 	if len(os.Args[1]) > 19 {
// 		return
// 	}
// 	switch os.Args[2] {
// 	case "+":
// 		fmt.Printf("%d\n", a+b)
// 	case "-":
// 		fmt.Printf("%d\n", a-b)
// 	case "*":
// 		fmt.Printf("%d\n", a*b)
// 	case "/":
// 		if b == 0 {
// 			fmt.Printf("No division by 0\n")
// 			return
// 		}
// 		fmt.Printf("%d\n", a/b)
// 	case "%":
// 		if b == 0 {
// 			fmt.Printf("No modulo by 0\n")
// 			return
// 		}
// 		fmt.Printf("%d\n", a%b)
// 	}
// }

// func Atoi(s string) int {
// 	neg := 1
// 	if s[0] == '-' {
// 		neg = -1
// 		s = s[1:]
// 	}
// 	n := 0
// 	for _, i := range s {
// 		if i < '0' || i > '9' {
// 			os.Exit(0)
// 		}
// 		n = n*10 + int(i-'0')
// 	}
// 	return neg * n
// }
// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	a := Atoi(os.Args[1])
// 	b := Atoi(os.Args[3])
// 	// if len(os.Args[1]) > 19 {
// 	// 	return
// 	// }
// 	// if a > 0 && b > 0 && a+b < 0 || a < 0 && b < 0 && a+b > 0 {
// 	// 	return
// 	}
// 	switch os.Args[2] {
// 	case "+":
// 		fmt.Printf("%d\n", a+b)
// 	case "-":
// 		fmt.Printf("%d\n", a-b)
// 	case "*":
// 		fmt.Printf("%d\n", a*b)
// 	case "/":
// 		if b == 0 {
// 			fmt.Printf("No division by 0\n")
// 			return
// 		} else {
// 			fmt.Printf("%d\n", a/b)
// 		}
// 	case "%":
// 		if b == 0 {
// 			fmt.Printf("No modulo by 0\n")
// 			return
// 		} else {
// 			fmt.Printf("%d\n", a%b)
// 		}
// 	}
// }

// func Atoi(s string) int {
// 	neg := 1
// 	if s[0] == '-' {
// 		neg = -1
// 		s = s[1:]
// 	}
// 	n := 0
// 	for _, i := range s {
// 		if i < '0' || i > '9' {
// 			os.Exit(0)
// 		}
// 		n = n*10 + int(i-'0')
// 	}
// 	return neg * n
// }
func main() {
	if len(os.Args) == 4 {
		a := Atoi(os.Args[1])
		b := Atoi(os.Args[3])
		if len(os.Args[1]) > 19 || len(os.Args[3]) > 19 {
			return
		}
		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
			return
		}
		switch os.Args[2] {
		case "+":
			fmt.Printf("%d\n", a+b)
		case "-":
			fmt.Printf("%d\n", a-b)
		case "*":
			fmt.Printf("%d\n", a*b)
		case "/":
			if b == 0 {
				fmt.Printf("No division by 0\n")
			} else {
				fmt.Printf("%d\n", a/b)
			}
		case "%":
			if b == 0 {
				fmt.Printf("No modulo by 0\n")
			} else {
				fmt.Printf("%d\n", a%b)
			}
		}
	}
}

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
