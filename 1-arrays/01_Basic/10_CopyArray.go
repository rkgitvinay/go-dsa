/*
Problem: Create a Copy of an Array

Description:
You are given an array of integers. Write a function `copyArray(arr []int) []int` that creates a copy of the given array and returns the new array.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].

Output:
- A new array that is a copy of the input array.

Example:

Input:
arr := []int{1, 2, 3, 4, 5}

Output:
[]int{1, 2, 3, 4, 5}

Constraints:
- The original array should not be modified.
- The solution should work in linear time (O(n)).
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	copy := copyArray(arr)
	fmt.Println("Copy array", copy)
}

func copyArray(arr []int) []int {
	copy := []int{}
	for _, val := range arr {
		copy = append(copy, val)
	}
	return copy
}
