// Find two numbers in a sorted array that sum up to a given target.
// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1] (indices of the two numbers that add up to 9)
package main

import "fmt"

func twoSum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{arr[left], arr[right]}
		} else if sum < target {
			right++
		} else {
			right--
		}
	}

	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(arr, target)
	fmt.Println("Result", result)
}
