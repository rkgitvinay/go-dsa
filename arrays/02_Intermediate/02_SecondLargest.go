/*
Problem: Find Second Largest Element

Description:
You are given an array of integers. Write a function `secondLargest(arr []int) int` that finds and returns the second largest element in the array. If the array has fewer than two elements, return -1.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].

Output:
- The second largest element in the array, or -1 if the array has fewer than two elements.

Example:

Input:
arr := []int{1, 2, 3, 4, 5}

Output:
5

Input:
arr := []int{1, 1, 1, 1}

Output:
-1

Constraints:
- Do not sort the array.
- The solution should run in linear time (O(n)).
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 11, 10, 9}
	secondLargestEle := secondLargest(arr)
	fmt.Println("second largest ele", secondLargestEle)
}

func secondLargest(arr []int) int {
	largest := 0
	secondLargest := -1

	for _, val := range arr {
		if val > largest {
			secondLargest = largest
			largest = val
		} else if val > secondLargest && val < largest {
			secondLargest = val
		}
	}
	return secondLargest
}
