/*
Problem: Sum and Average of Array Elements

Description:
You are given an array of integers. Your task is to compute the sum and the average of the elements in the array.
You need to write a function `sumAndAverage(arr []int) (int, float64)` that takes an array of integers and returns the sum
of the elements and the average of the elements as a tuple `(sum, average)`.

Input:
- An array arr of integers where 1 <= len(arr) <= 10^5 and -10^3 <= arr[i] <= 10^3 for each element arr[i].

Output:
- A tuple (sum, average) where:
  - sum is the total sum of the array elements.
  - average is the average of the array elements, calculated as sum divided by the number of elements.

Example:

Input:
arr := []int{4, 2, 9, 7, 3, 1, 5}

Output:
(31, 4.428571428571429)

Constraints:
- The function should solve the problem in linear time (O(n)), where n is the number of elements in the array.
- The solution should not use any additional space beyond the output.
*/

package main

import "fmt"

func main() {
	arr := []int{4, 5, 1, 6, 7, 2}

	sum, avg := sumAndAverage(arr)
	fmt.Printf("Sum: %d, Avg: %f\n", sum, avg)
}

func sumAndAverage(arr []int) (int, float32) {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	avg := float32(sum) / float32(len(arr))

	return sum, float32(avg)
}
