package main

// I pick a number from 1 to n. You have to guess which number I picked.

// Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

// You call a pre-defined API int guess(int num), which returns three possible results:
// -1 if num > pick ; 1 if num < pick ; 0 if num == pick

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		mid := low + (high-low)/2
		if guess(mid) == -1 {
			high = mid - 1
		} else if guess(mid) == 1 {
			low = mid + 1
		} else if guess(mid) == 0 {
			return mid
		}
	}
	return 1
}

func main() {

}
