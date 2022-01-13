package main

import "fmt"

// **Algorithm**
// - initialize left to equal 0
// - initialize right to equal length of array - 1
// - while left <= right
//   - initialize mid to equal left + ((right - left)/2)
//   - if value to right of mid is less than mid, return value to right of mid
//   - if last element is less than value at mid
//     - set left to mid + 1
//   - otherwise, set right to mid - 1

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	// if nums is only 1 element, return element
	if len(nums) == 1 {
		return nums[0]
	}

	for left <= right {
		mid := left + ((right - left) / 2)

		// guard against out of bounds when mid is first or last element
		if mid == 0 {
			if nums[mid+1] < nums[mid] {
				return nums[mid+1]
			} else {
				return nums[0]
			}
		} else if mid == len(nums)-1 {
			if nums[mid-1] > nums[mid] {
				return nums[mid]
			}
		} else if nums[mid+1] < nums[mid] {
			return nums[mid+1]
		} else if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		// check value of first and last element
		if nums[right] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))          // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))    // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))         // 11
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))          // 1
	fmt.Println(findMin([]int{1, 2}))                   // 1
	fmt.Println(findMin([]int{1, 2, 3, 4}))             // 1
	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))          // 1
	fmt.Println(findMin([]int{2, 3, 4, 5, 6, 7, 8, 0})) // 2
	fmt.Println(findMin([]int{3, 4, 5, 6, 1, 2}))       // 1
}
