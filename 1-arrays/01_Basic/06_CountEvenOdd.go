/*
Problem: Count the Number of Even and Odd Elements

Description:
You are given an array of integers. Write a function `countEvenOdd(arr []int) (int, int)` that counts the number of even and odd elements in the array and returns the counts.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].

Output:
- Two integers:
  1. The count of even elements in the array.
  2. The count of odd elements in the array.

Example:

Input:
arr := []int{1, 2, 3, 4, 5, 6}

Output:
Even Count: 3, Odd Count: 3

Constraints:
- The solution should run in linear time (O(n)), where n is the length of the array.
- Solve the problem using a single traversal of the array.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 6}
	even, odd := countEvenOdd(arr)
	fmt.Printf("Event Count: %d, Odd Count: %d\n", even, odd)
}

func countEvenOdd(arr []int) (int, int) {
	even, odd := 0, 0

	for _, val := range arr {
		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}
