// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) == 2 {
// 		n := 0
// 		for _, i := range os.Args[1] {
// 			if i < '0' || i > '9' {
// 				fmt.Printf("ERROR: cannot convert to roman digit\n")
// 				return
// 			}
// 			n = n*10 + int(i-'0')
// 		}
// 		if n == 0 || n > 4000 {
// 			fmt.Printf("ERROR: cannot convert to roman digit\n")
// 			return
// 		}
// 		num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
// 		romnum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
// 		mathnum := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
// 		var str, sum string
// 		for v := range num {
// 			temp := n / num[v]
// 			if temp > 0 {
// 				n -= num[v]
// 				str += romnum[v]
// 				sum += mathnum[v] + "+"
// 				temp--
// 			}
// 		}
// 		fmt.Printf("%s\n%s\n", sum[:len(sum)-1], str)
// 	}
// }

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	a := Atoi(os.Args[1])
	if a <= 0 || a >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit")
		fmt.Printf("\n")
		return
	}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hnds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thns := []string{"", "M", "MM", "MMM"}
	res := thns[a/1000] + hnds[a%1000/100] + tens[a%100/10] + ones[a%10]
	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
	hnds1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
	thns1 := []string{"", "M+", "M+M+", "M+M+M+"}
	res1 := thns1[a/1000] + hnds1[a%1000/100] + tens1[a%100/10] + ones1[a%10]
	fmt.Printf(res1[:len(res1)-1])
	fmt.Printf("\n")
	fmt.Printf(res)
	fmt.Printf("\n")
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
			return 0
		}
		n = n*10 + int(i-'0')
	}
	return neg * n
}
