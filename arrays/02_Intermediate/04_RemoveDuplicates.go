/*
Problem: Remove Duplicates from Sorted Array

Description:
You are given a sorted array of integers. Write a function `removeDuplicates(arr []int) []int` that removes the duplicate elements in place and returns the updated array containing only unique elements.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5.

Output:
- A slice containing only the unique elements in their original order.

Examples:

Input:
arr := []int{1, 1, 2, 2, 3, 4, 4, 5}

Output:
[1, 2, 3, 4, 5]

Input:
arr := []int{1, 1, 1, 1, 1}

Output:
[1]

Constraints:
- The function should run in O(n) time.
- Do not use any additional data structures.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
	arr = removeDuplicates(arr)
	fmt.Println(arr)
}

func removeDuplicates(arr []int) []int {
	unique := []int{}
	seen := make(map[int]bool)

	for _, val := range arr {
		if !seen[val] {
			unique = append(unique, val)
			seen[val] = true
		}
	}
	return unique
}
