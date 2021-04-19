package main

import (
	"fmt"
	"homework/serebryakov"
)

func main() {
	fmt.Println(serebryakov.SolutionSquareGenerator(9, 5))
	fmt.Println(serebryakov.SolutionSquareGenerator(1, 10))
	fmt.Println(serebryakov.SolutionSquareGenerator(2, 1))
	fmt.Println(serebryakov.SolutionSquareGenerator(5, 0))
}
