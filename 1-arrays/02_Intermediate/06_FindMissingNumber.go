/*
Problem: Find Missing Number in a Series

Given an array containing distinct integers from `1` to `n` with one number missing, find and return the missing number.

The array contains `n-1` numbers, and you need to identify the number from `1` to `n` that is absent.

Input:
- An array `arr` of size `n-1` where `1 <= len(arr) < n`
- An integer `n` representing the total size of the series from `1` to `n`.

Output:
- An integer representing the missing number in the series.

Examples:

Example 1:
Input: arr = [1, 2, 3, 5], n = 5
Output: 4

Example 2:
Input: arr = [7, 8, 9, 10, 11], n = 12
Output: 12

Example 3:
Input: arr = [1, 3], n = 3
Output: 2

Constraints:
- The array `arr` contains distinct integers and is unsorted.
- The array size is `n-1` and the number from 1 to n is missing.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5}
	missing := findMissingNumber(arr)
	fmt.Println("missing", missing)
}

func findMissingNumber(arr []int) int {
	n := len(arr) + 1
	expectedSum := (n * (n + 1)) / 2
	actualSum := 0

	for _, val := range arr {
		actualSum += val
	}

	return expectedSum - actualSum
}
