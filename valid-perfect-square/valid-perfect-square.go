package main

import "fmt"

// **Algorithm**
// - intialize minimum to equal 1
// - intialize maximum to equal input number
// - while minimum <= maximum
//   - initialize midpoint variable and set equal to minimum + ((maximum - minimum) / 2)
//   - find square of midpoint
//   - if square of midpoint == input
//     - return true
//   - if square of midpoint is < input
//     - minimum = midpoint + 1
//   - if square of midpoint is > input
//     - maximum = midpoint - 1
// - return false if exited loop without returning true

func isPerfectSquare(num int) bool {
	min, max := 1, num

	for min <= max {
		mid := min + ((max - min) / 2)
		square := mid * mid

		if square == num {
			return true
		} else if square < num {
			min = mid + 1
		} else if square > num {
			max = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16)) // true
	fmt.Println(isPerfectSquare(14)) // false
}
