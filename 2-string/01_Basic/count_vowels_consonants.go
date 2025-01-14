/*
Problem: Count Vowels and Consonants

Description:
Given a string `s`, your task is to count the number of vowels and consonants in the string. Both uppercase and lowercase letters should be considered.

Write a function countVowelsAndConsonants(s string) (int, int) that returns two values:
- The count of vowels.
- The count of consonants.

Input:
- A string `s` where 1 <= len(s) <= 10^5.

Output:
- Two integers:
  1. The number of vowels in the string.
  2. The number of consonants in the string.

Example 1:

Input:
s := "hello"

Output:
2, 3

Example 2:

Input:
s := "Beautiful"

Output:
5, 4

Example 3:

Input:
s := "12345!"

Output:
0, 0

Constraints:
- Only alphabetic characters should be considered (ignore digits, symbols, etc.).
- Treat 'a', 'e', 'i', 'o', 'u' (both uppercase and lowercase) as vowels.
- The function should solve the problem in O(n) time complexity.
*/

package main

import (
	"fmt"
	"unicode"
)

// count_vowels_consonants: Implement your solution here
func countVowelsAndConsonants(s string) (int32, int32) {
	vowelCount, consonantCount := 0, 0

	for _, char := range s {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if isVowel(lowerChar) {
				vowelCount++
			} else {
				consonantCount++
			}

		}
	}
	return int32(vowelCount), int32(consonantCount)
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

func main() {
	// Example 1
	s1 := "hello"
	v1, c1 := countVowelsAndConsonants(s1)
	fmt.Printf("String: '%s' | Vowels: %d, Consonants: %d\n", s1, v1, c1)

	// Example 2
	s2 := "Beautiful"
	v2, c2 := countVowelsAndConsonants(s2)
	fmt.Printf("String: '%s' | Vowels: %d, Consonants: %d\n", s2, v2, c2)
}
