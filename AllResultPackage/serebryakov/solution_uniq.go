package serebryakov

//Подсчет количества уникальных значений в массиве А
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
