package main

import "fmt"

// **Algorithm**
//   - initialize left variable to 0
//   - initialize right variable to length of array - 1
//   - initialize mid to right - ((right - left) / 2)
//   - while left <= right
//     - if values to both left and right of mid are less than mid, return mid
//     - if value to right of mid is greater than mid
//       - set left to equal mid + 1
//     - else set right to equal mid - 1

func isMidPeak(mid int, nums []int) bool {
	// guard against out of bounds by checking if mid is first or last element
	if mid == 0 {
		if nums[mid+1] < nums[mid] {
			return true
		}
	} else if mid == len(nums)-1 {
		if nums[mid-1] < nums[mid] {
			return true
		}
	} else if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
		return true
	}
	return false
}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	// if nums is only one element return index 0
	if len(nums) == 1 {
		return 0
	}

	for left <= right {
		mid := left + ((right - left) / 2)

		if isMidPeak(mid, nums) {
			return mid
		}

		if nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          //	index 2
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) //	index 1 or 5
	fmt.Println(findPeakElement([]int{2, 1}))                //	index 0
	fmt.Println(findPeakElement([]int{1}))                   //	index 0
	fmt.Println(findPeakElement([]int{1, 2}))                //	index 1
}
