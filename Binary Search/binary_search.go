package main

import (
	"fmt"
)

// binarySearch finds the target in an array in
// O(log(n)) runtime
func binarySearch(arr []int, target int) int {
	start := 0               // start will be used to move in an array
	end := len(arr) - 1      // end will be used to point the end of the array
	mid := (start + end) / 2 // mid will help finding the mid index of array and the target value
	// Move in an array from start to end
	for start <= end {
		if arr[mid] < target {
			// If arr[mid] is smaller than target
			// move start to right side of the array
			start = mid + 1
		}
		if arr[mid] > target {
			// If arr[mid] is bigger than target
			// move end to left side of the array
			end = mid - 1
		}
		if arr[mid] == target {
			// Return mid if arr[mid] is equal to target
			return mid
		}
		// Find the mid index again
		mid = (start + end) / 2
	}
	return -1 // Return -1 if the target value not found
}

func main() {
	target := 8
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := binarySearch(arr, target)
	fmt.Println(res)
}
