package array_slices

func Sum(numbers []int) int {
	var sum int
	//range returns two values - the index and the value
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arraysToSum ...[]int) []int {
	var sums []int
	for _, arrayWithInt := range arraysToSum {
		sums = append(sums, Sum(arrayWithInt))
	}

	return sums
}
