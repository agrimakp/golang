package main

func average(salary []int) float64 {
	minNumber := 1000000
	maxNumber := 1000
	total := 0
	count := float64(len(salary) - 2)

	for _, s := range salary {
		total += s
		if s < minNumber {
			minNumber = s
		}
		if s > maxNumber {
			maxNumber = s
		}

	}
	return float64(total-minNumber-maxNumber) / float64(count)
}

func main() {
	var salary = []int{4000, 3000, 1000, 2000}
	var n = average()
}

// Given two non-negative integers low and high.
// Return the count of odd numbers between low and high (inclusive).
