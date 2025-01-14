/*
Problem: Valid Anagram

Description:
Given two strings, your task is to determine if one string is an anagram of the other.
An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase,
using all the original letters exactly once.

Write a function isAnagram(s string, t string) bool that returns true if `t` is an anagram of `s`, and false otherwise.

Input:
- Two strings `s` and `t` where 1 <= len(s), len(t) <= 5 * 10^4.
- Both strings consist only of lowercase English letters.

Output:
- A boolean value `true` if `t` is an anagram of `s`, otherwise `false`.

Example 1:

Input:
s := "anagram"
t := "nagaram"

Output:
true

Example 2:

Input:
s := "rat"
t := "car"

Output:
false

Constraints:
- The function should solve the problem in O(n) time complexity, where n is the length of the longer string.
- The solution should use O(1) additional space if we assume the alphabet size is fixed (26 lowercase letters).

Follow-up:
- How would you handle the case where the strings contain Unicode characters instead of just lowercase English letters?
*/

package main

import "fmt"

// check_anagram: Implement your solution here
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// create frequency map of characters in s
	charCount := make(map[rune]int64)
	for _, char := range s {
		charCount[char]++
	}

	for _, ch := range t {
		if charCount[ch] == 0 {
			return false
		}
		charCount[ch]--
	}

	return true
}

func main() {
	fmt.Println("Running: check_anagram")

	s := "anagram"
	t := "nagaram"
	check := isAnagram(s, t)
	fmt.Println("Is Anagram ", check)

	s2 := "rat"
	t2 := "car"
	check2 := isAnagram(s2, t2)
	fmt.Println("Is Anagram ", check2)

	s3 := "aacc"
	t3 := "ccac"
	check3 := isAnagram(s3, t3)
	fmt.Println("Is Anagram ", check3)
}
