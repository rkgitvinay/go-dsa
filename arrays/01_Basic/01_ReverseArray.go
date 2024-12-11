/*
Problem: Reverse an Array

Description:
Given an array of integers, your task is to reverse the array. You need to write a function reverseArray(arr []int) []int
that takes an array of integers and returns a new array where the elements are in the reverse order.

Input:
- An array arr of integers where 1 <= len(arr) <= 10^5 and -10^3 <= arr[i] <= 10^3 for each element arr[i].

Output:
- A new array where the elements are in reverse order compared to the original array.

Example:

Input:
arr := []int{1, 2, 3, 4, 5}

Output:
[]int{5, 4, 3, 2, 1}

Constraints:
- The function should solve the problem in linear time (O(n)), where n is the number of elements in the array.
- The solution should not use any extra space, meaning it should modify the array in-place if possible.
*/
package main

import "fmt"

func reverseArray(arr []int) []int {
	var rev []int
	var size = len(arr)

	for i := size - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}

	return rev
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reversedArr := reverseArray(arr)
	fmt.Println("Reversed Array: ", reversedArr)
}
