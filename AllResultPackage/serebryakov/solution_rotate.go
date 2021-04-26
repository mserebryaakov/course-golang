package serebryakov

//Перестановка left и right элементов массива А, пока left < right
func ReverseElements(A []int, left int, right int) {
	var tmp int
	for left < right {
		tmp = A[left]
		A[left] = A[right]
		A[right] = tmp
		left++
		right--
	}
}

//Сдвиг массива А вправо (k раз)
func SolutionRotate(A []int, k int) []int {
	ReverseElements(A, 0, len(A)-1)
	ReverseElements(A, 0, k-1)
	ReverseElements(A, k, len(A)-1)
	return A
}
