/*
Problem: Remove an Element from an Array

Description:
You are given an array of integers and a target element. Write a function `removeElement(arr []int, target int) []int` that removes all occurrences of the target element from the array and returns the updated array.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].
- An integer `target` representing the element to remove.

Output:
- An array with all occurrences of the target element removed.

Example:

Input:
arr := []int{1, 2, 3, 4, 2, 5, 2}
target := 2

Output:
[]int{1, 3, 4, 5}

Constraints:
- The solution should run in linear time (O(n)).
- Do not use additional arrays if possible.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 2, 5, 2}
	target := 2

	result := removeElement(arr, target)
	fmt.Println(result)
}

func removeElement(arr []int, target int) []int {
	result := []int{}
	for _, val := range arr {
		if val != target {
			result = append(result, val)
		}
	}
	return result
}
