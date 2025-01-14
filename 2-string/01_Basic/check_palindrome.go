/*
Problem: Valid Palindrome

Description:
Given a string, your task is to determine if it is a palindrome. A string is a **palindrome** if it reads the same backward as forward, ignoring case, spaces, and non-alphanumeric characters.

Write a function isPalindrome(s string) bool that returns true if the string is a palindrome, and false otherwise.

Input:
- A string `s` where 1 <= len(s) <= 2 * 10^5.

Output:
- A boolean value `true` if the string is a palindrome, otherwise `false`.

Example 1:

Input:
s := "A man, a plan, a canal: Panama"

Output:
true

Example 2:

Input:
s := "race a car"

Output:
false

Constraints:
- Only alphanumeric characters should be considered.
- The function should ignore case differences.
- Solve the problem with O(n) time complexity.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// check_palindrome: Implement your solution here
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	fmt.Printf("Is '%s' a palindrome? %v\n", s1, isPalindrome(s1)) // true

	// Example 2
	s2 := "race a car"
	fmt.Printf("Is '%s' a palindrome? %v\n", s2, isPalindrome(s2)) // false
}
