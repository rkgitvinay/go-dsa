/*
Problem: Find Maximum and Minimum Element in an Array

Description:
You are given an array of integers. Your task is to find both the maximum and the minimum elements in the array.
You need to write a function findMaxMin(arr []int) (int, int) that takes an array of integers and returns a tuple containing
the maximum and minimum elements in the array.

Input:
- An array arr of integers where 1 <= len(arr) <= 10^5 and -10^3 <= arr[i] <= 10^3 for each element arr[i].

Output:
- A tuple (max, min) where max is the maximum element in the array and min is the minimum element in the array.

Example:

Input:
arr := []int{4, 2, 9, 7, 3, 1, 5}

Output:
(9, 1)

Constraints:
- The function should solve the problem in linear time (O(n)), where n is the number of elements in the array.
- The solution should not use any additional space beyond the output.
*/

package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 7, 2}
	max, min := findMaxMin(arr)
	fmt.Printf("Maximum: %d, Minimum: %d\n", max, min)
}

func findMaxMin(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}
