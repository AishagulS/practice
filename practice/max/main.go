// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{-23, -123454435555554554, 1, 1125245554, 55, 935340001000000202}
// 	max := Max(a)
// 	fmt.Println(max)
// }

// func Max(a []int) int {
// 	var max int
// 	for _, i := range a {
// 		if i > max {
// 			max = i
// 		}
// 	}
// 	return max
// }
package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	var max int
	for i := range a {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
