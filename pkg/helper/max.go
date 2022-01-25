package helper

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func MaxIndex(numbers []int) int {
	var max int
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}
	}

	return index
}
