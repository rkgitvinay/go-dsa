/*
Problem: Check if Array is Sorted

Description:
You are given an array of integers. Your task is to check if the array is sorted in non-decreasing order.
You need to write a function `isArraySorted(arr []int) bool` that takes an array of integers and returns a boolean value:
- `true` if the array is sorted in non-decreasing order (i.e., each element is less than or equal to the next element).
- `false` if the array is not sorted.

Input:
- An array arr of integers where 1 <= len(arr) <= 10^5 and -10^3 <= arr[i] <= 10^3 for each element arr[i].

Output:
- A boolean value `true` if the array is sorted, otherwise `false`.

Example:

Input:
arr := []int{1, 2, 3, 4, 5}

Output:
true

Input:
arr := []int{5, 3, 4, 2, 1}

Output:
false

Constraints:
- The function should solve the problem in linear time (O(n)), where n is the number of elements in the array.
- The solution should not use any additional space beyond the output.
*/

package main

import "fmt"

func main() {
	arr1 := []int{4, 5, 5, 5, 6}
	arr2 := []int{4, 5, 2, 5, 6}

	isArray1Sorted := isArraySorted(arr1)
	isArray2Sorted := isArraySorted(arr2)

	fmt.Println("Is array 1 sorted: ", isArray1Sorted)
	fmt.Println("Is array 2 sorted: ", isArray2Sorted)
}

func isArraySorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
