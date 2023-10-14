// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	a := os.Args[1:]
// 	for _, i := range os.Args[len(a)] {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	b := os.Args[len(os.Args)-1]
	for _, s := range b {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
