func minSUbArrayLen(target int, nums []int) int { 
	a, r := 0, 0

	result := len(nums) + 1
	currentSum := 0

	for r < len(nums) {
		currentSum += nums[r]

		for currentSum >= target {
			currentLength := r - a + 1
			if currentLength == 1 {
				return currentLength
			}
			if currentLength < result {
				result = currentLength
			}
			currentSum -= nums[a]
			a++
		}
		r++
	}

	if (result == lens(nums) + 1) {
		return 0
	}
}
	// brute force approach?

	//two nested for loops
	// inside of it you are calculation sum? [2, 3, 1] calcualating sum O(N)
		// if sum is >= target checking length? reassigning result if it is less
	// Time complexity? O(N^3)

	// Two nexted for O(N^2)
		// don't recalculate sum for entire slice each time
		// instead for sum: 
			// when you move i, sum -= i
			// when you move j, sum += j
			
// Possible to solve with Binary Search O(N log N)


// Pointer solution: The O(N) solution
	// 1. Where do my pointers start?
				// Anchor and Runner both start at 0
				// Sum of values between pointers variable
	// 2. When do I move the Runner pointer?
				// When sum is less than target
	// 3. When do i move the Anchor pointer?
				// when sum is greater and equal to target
	// 4. Do I do something else?
				// when a anchor moved, first subtract anchor value from sum
				// after runner moved, add the runner value to sum
				// length = runner - anchor + 1
	// 5. When do we end?
				// when result == 1
				// when Runner is >= length of array
/*  [2, 3, 1,2, 4, 3]  target = 7, sum = 2
	   a
		 r	

		 [2, 3, 1,2, 4, 3]  taget = 7, sum = 8, length = 4, result = 4
	   a
		 					r	

		[2, 3, 1,2, 4, 3]  taget = 7, sum = 7, length = 3, result = 3
	      	 a
		 					  r	

		[2, 3, 1,2, 4, 3]  taget = 7, sum = 7, length = 2, result = 2
	      	      a
		 					  	 r
										
										
