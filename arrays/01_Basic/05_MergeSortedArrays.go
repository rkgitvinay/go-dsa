/*
Problem: Merge Two Sorted Arrays

Description:
You are given two sorted arrays of integers. Your task is to merge them into a single sorted array.
You need to write a function `mergeSortedArrays(arr1 []int, arr2 []int) []int` that takes two sorted arrays and returns a new array
containing all elements from both arrays, sorted in non-decreasing order.

Input:
- Two sorted arrays `arr1` and `arr2` where 1 <= len(arr1), len(arr2) <= 10^5 and -10^3 <= arr1[i], arr2[i] <= 10^3 for each element arr1[i] and arr2[i].

Output:
- A new array containing all elements from both arrays, sorted in non-decreasing order.

Example:

Input:
arr1 := []int{1, 3, 5, 7}
arr2 := []int{2, 4, 6, 8}

Output:
[]int{1, 2, 3, 4, 5, 6, 7, 8}

Constraints:
- The function should solve the problem in linear time (O(n + m)), where n and m are the lengths of the two arrays.
- The solution should not use any additional space beyond the output array.
*/
func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}

	mergedArr := mergeSortedArrays(arr1, arr2)
	fmt.Println("Merged Array:", mergedArr)
}

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	merged := []int{}

}