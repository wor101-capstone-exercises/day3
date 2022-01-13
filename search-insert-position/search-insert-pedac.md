
**Problem**
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Constraints:

  1 <= nums.length <= 104
  -104 <= nums[i] <= 104
  nums contains distinct values sorted in ascending order.
  -104 <= target <= 104

Understanding the Problem

- input:
  - array of integers sorted ascending
  - an integer that is the target number

- output:
  - integer
  - either index of where the target number is in the array
  - or index of where target should be in the array


- model of problem O(log N)
  - perform binary search on array starting with middle index


**Examples / Test Cases**
Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

**Data Structures**


**Algorithm**
