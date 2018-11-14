package slices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	numberOfSlices := len(numbersToSum)
	sums := make([]int, numberOfSlices)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}