package serebryakov

//Генератор n квадратов натуральных чисел, начиная с start
func SolutionSquareGenerator(start int, n int) []int {
	if start <= 0 {
		n += start - 1
		start = 1
	}

	var result []int = make([]int, n)
	for i := 1; i <= n; i++ {
		result[i-1] = start * start
		start++
	}
	return result
}
