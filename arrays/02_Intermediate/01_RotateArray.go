/*
Problem: Rotate an Array

Description:
You are given an array of integers and an integer `k`. Write a function `rotateArray(arr []int, k int) []int` that rotates the array to the right by `k` steps. If `k` is greater than the length of the array, perform `k % len(arr)` rotations.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].
- An integer `k` representing the number of steps to rotate.

Output:
- A new array that is the result of rotating the input array by `k` steps.

Example:

Input:
arr := []int{1, 2, 3, 4, 5}
k := 2

Output:
[]int{4, 5, 1, 2, 3}

Input:
arr := []int{1, 2, 3, 4, 5}
k := 7

Output:
[]int{4, 5, 1, 2, 3}

Constraints:
- Solve the problem in linear time (O(n)).
- Use only O(1) extra space if possible.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2

	rotated := rotateArray(arr, k)
	fmt.Println("Rotated array", rotated)
}

func rotateArray(arr []int, k int) []int {
	rotated := make([]int, len(arr))
	n := len(arr)
	rotated = append(arr[n-k:], arr[:n-k]...)
	return rotated
}
