
**Problem**
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.

Understanding the Problem

- input:
  - an integer

- output:
  - a boolean


- model of problem
  - do a binary search on possible solutions
  - start with 1 as minimum(left) and start with input number as maximum(right)
  - find the midpoint(possible solution)
  - if midpoint * midpoint equals input, return true
  - if midpoint * midpoint > input, go right
  - if midpoint * midpoint < input, go left


**Examples / Test Cases**


**Data Structures**


**Algorithm**
- intialize minimum to equal 1
- intialize maximum to equal input number
- while minimum <= maximum
  - initialize midpoint variable and set equal to minimum + ((maximum - minimum) / 2)
  - find square of midpoint
    - if square of midpoint == input
      - return true
    - if square of midpoint is < input
      - minimum = midpoint + 1
    - if square of midpoint is > input
      - maximum = midpoint - 1
- return false if exited loop without returning true
