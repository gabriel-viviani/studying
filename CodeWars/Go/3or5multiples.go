package kata

func Multiple3And5(number int) int {
	numbers := make([]int, number)
	var sum int
	for i := range numbers {
		sum += AddIfMultiple(i)
	}
	return sum
}

func AddIfMultiple(number int) int {
	var sum int
	if number%3 == 0 || number%5 == 0 {
		sum += number
	}
	return sum
}
