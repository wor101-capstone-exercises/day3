package main

func isValidLength(possSol, target int, nums []int) bool {
	a, r := 0, 0

	sum := 0

	for r < len(nums) {
		if r < possSol {
			sum += nums[r]
			r++
		} else {
			sum -= nums[a]
			a++
			sum += nums[r]
			r++
		}
		if sum >= target {
			return true
		}
	}
	return false
}

func minSubArrayLen(target int, nums []int) int {
	// O(N logN)

	// use mid point as possible result value, and do a binary search over the potential result values
	// mid point
	// what can i do binary search on?
	// binary search to check lengths

	// 1 (minimum) and 6 (maximum) (minimumLength=left + maximumLength=right) / 2

	// mid 7 / 2 = 3 -> current Length = 3

	// O(N) (try to find sum >= target of length 3)
	// can use sliding window to find sum with an O(N) over total array
	// new binary search 1 and 3

	min, max := 1, len(nums)
	result := 0

	for min <= max {
		possSol := min + ((max - min) / 2)
		if isValidLength(possSol, target, nums) { // use k-slide window
			result = possSol
			max = possSol - 1
		} else {
			min = possSol + 1
		}
	}
	return result
}

// [2,3,1,2,4,3], target = 7
