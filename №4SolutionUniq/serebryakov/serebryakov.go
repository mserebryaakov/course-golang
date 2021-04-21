package serebryakov

func SolutionUniq(A []int) int {
	mapNumbers := make(map[int]int)
	var resultCount int = len(A)
	for _, num := range A {
		mapNumbers[num] = mapNumbers[num] + 1
		if mapNumbers[num] == 2 {
			resultCount -= 2
		}
		if mapNumbers[num] > 2 {
			resultCount--
		}
	}
	return resultCount
}

func SolutionUniq2(A []int) int {
	mapNumbers := make(map[int]int)
	var resultCount int = 0
	for _, num := range A {
		mapNumbers[num] = mapNumbers[num] + 1
	}
	for _, num := range A {
		if mapNumbers[num] == 1 {
			resultCount++
		}
	}
	return resultCount
}
