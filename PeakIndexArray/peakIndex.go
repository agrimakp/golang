package main

func peakIndexInMountainArray(arr []int) int {
	low := 0
	high := len(arr) - 1
	var mid int

	for low < high {
		mid = low + (high-low)/2
		guess := arr[mid]
		right := arr[mid+1]
		if guess > right {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}

func main() {

}
