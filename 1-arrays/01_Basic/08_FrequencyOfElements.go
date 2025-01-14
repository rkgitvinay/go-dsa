/*
Problem: Find the Frequency of Each Element

Description:
You are given an array of integers. Write a function `findFrequency(arr []int) map[int]int` that calculates the frequency of each element in the array and returns the result as a map.

Input:
- An array `arr` of integers, where 1 <= len(arr) <= 10^5 and -10^5 <= arr[i] <= 10^5 for each element arr[i].

Output:
- A map where the keys are the unique elements of the array and the values are their corresponding frequencies.

Example:

Input:
arr := []int{1, 2, 2, 3, 3, 3, 4}

Output:
{
  1: 1,
  2: 2,
  3: 3,
  4: 1
}

Constraints:
- Solve the problem in linear time (O(n)).
- Use a map to store the frequencies.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 3, 5, 4}
	frequencies := findFrequency(arr)

	for key, val := range frequencies {
		fmt.Printf("Key %d occured %d times\n", key, val)
	}
}

func findFrequency(arr []int) map[int]int {
	frequencies := make(map[int]int)

	for _, val := range arr {
		frequencies[val]++
	}

	return frequencies
}
