package serebryakov

func SolutionSquareGenerator(start int, n int) []int {
	var result []int = make([]int, n)
	for i := 1; i <= n; i++ {
		result[i-1] = start * start
		start++
	}
	return result
}