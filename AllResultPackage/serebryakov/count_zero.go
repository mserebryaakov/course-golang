package serebryakov

//Поиск самой длинной последовательности нулей в двоичном представлении числа N
func SolutionBinaryGap(N int) int {
	var count int = 0
	var result int = 0
	for N != 0 {
		if N%2 == 0 {
			count++
		} else {
			if count > result {
				result = count
			}
			count = 0
		}
		N = N >> 1
	}
	return result
}
