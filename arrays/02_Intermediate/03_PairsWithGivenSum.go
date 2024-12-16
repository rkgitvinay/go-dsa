/*
Problem: Find All Pairs with a Given Sum

Description:
You are given an array of integers and a target sum. Write a function `findPairs(arr []int, target int) [][]int` that finds all unique pairs of elements in the array whose sum equals the target.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5.
- An integer `target` representing the target sum.

Output:
- A 2D slice containing all pairs `[a, b]` where `a + b = target`.
- If no such pairs exist, return an empty slice.

Examples:

Input:
arr := []int{1, 2, 3, 4, 5}
target := 6

Output:
[[1, 5], [2, 4]]

Input:
arr := []int{3, 3, 3, 3}
target := 6

Output:
[[3, 3]]

Constraints:
- The solution should run in O(n) time for an optimal approach.
- The same pair cannot appear more than once in the output.
- Do not use the same element from the array twice for a pair.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 6
	pairs := findPairs(arr, target)
	fmt.Println("Pairs", pairs)
}

// NOTE - I will use simple and Brute force method
func findPairs(arr []int, target int) [][]int {
	pairs := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{arr[i], arr[j]})
			}
		}
	}
	return pairs
}

func findPairsUsingMap(arr []int, target int) [][]int {

}
