**Problem**
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

Constraints:
  n == nums.length
  1 <= n <= 5000
  -5000 <= nums[i] <= 5000
  All the integers of nums are unique.
  nums is sorted and rotated between 1 and n times.

Understanding the Problem

- input:
  - an array of nums
  - no duplicates
  - sorted in ascending, but rotated a minimum of 1 time

- output:
  - an integer representing the minimum number in the array

- model of problem Need O(log N)
  - find middle index of array
  - if value to right of middle is less than value at middle, it is the point of rotation
  - check first and last element of array O(1)
  - if the last element is less than value of middle index, than the rotation must be to the right
  - if the first element is greater than the value of the middle index, than the rotation must be to the left


**Examples / Test Cases**
Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times. 

**Data Structures**


**Algorithm**
- initialize left to equal 0
- initialize right to equal length of array - 1
- while left <= right
  - initialize mid to equal left + ((right - left)/2)
  - if value to right of mid is less than mid, return value to right of mid
  - if last element is less than value at mid
    - set left to mid + 1
  - otherwise, set right to mid - 1