package main

// Sum for numbers
func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
