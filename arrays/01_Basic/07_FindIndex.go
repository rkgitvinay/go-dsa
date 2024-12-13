/*
Problem: Find the Index of a Target Element

Description:
You are given an array of integers and a target element. Write a function `findIndex(arr []int, target int) int` that finds the index of the target element in the array. If the target element is not found, return -1.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].
- An integer `target` to search for in the array.

Output:
- An integer representing the index of the target element in the array. If the target element is not present, return -1.

Example:

Input:
arr := []int{5, 3, 7, 1, 9}
target := 7

Output:
2

Input:
arr := []int{5, 3, 7, 1, 9}
target := 4

Output:
-1

Constraints:
- Solve the problem in linear time (O(n)).
*/

package main

import "fmt"

func main() {
	arr := []int{5, 3, 7, 1, 9}
	target := 4
	index := findIndex(arr, target)
	if index != -1 {
		fmt.Printf("Found ele at indedx %d\n", index)
	} else {
		fmt.Println("Ele not found")
	}
}

func findIndex(arr []int, target int) int {
	index := -1
	for idx, val := range arr {
		if val == target {
			index = idx
			return index
		}
	}
	return index
}
