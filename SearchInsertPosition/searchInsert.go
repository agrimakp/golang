package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1

		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {

}
