package main

import "fmt"

func Chunk(slice []int, size int) {
	emptyslice := [][]int{}
	if len(slice) == 0 {
		fmt.Println(emptyslice)
		return
	}
	if size <= 0 {
		fmt.Println()
		return
	}
	for len(slice) > size {
		emptyslice = append(emptyslice, slice[:size])
		slice = slice[size:]
	}
	emptyslice = append(emptyslice, slice)
	fmt.Println(emptyslice)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
