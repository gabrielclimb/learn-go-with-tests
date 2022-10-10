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
	lengthOfNumbers := len(arraysToSum)
	// make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.
	sums := make([]int, lengthOfNumbers)

	for index, arraysWithInts := range arraysToSum {
		sums[index] = Sum(arraysWithInts)
	}
	return sums
}
