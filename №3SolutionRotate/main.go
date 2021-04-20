package main

import (
	"fmt"
	"homework/serebryakov"
)

func main() {
	for i := 1; i < 10; i++ {
		A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(serebryakov.SolutionRotate(A, i))
	}
}
