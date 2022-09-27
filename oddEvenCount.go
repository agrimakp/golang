package main

func countOdds(low int, high int) int {
	var count int = 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}
