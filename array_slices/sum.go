package array_slices

func Sum(numbers [5]int) int {
	var sum int
	//range returns two values - the index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}
