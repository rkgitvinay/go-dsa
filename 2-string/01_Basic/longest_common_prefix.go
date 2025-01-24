/*
Problem: Longest Common Prefix

Description:
Write a function longestCommonPrefix(strs []string) string that finds the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Input:
- An array of strings `strs` where 1 <= len(strs) <= 200 and 0 <= len(strs[i]) <= 200.

Output:
- A string representing the longest common prefix.

Example 1:

Input:
strs := []string{"flower", "flow", "flight"}

Output:
"fl"

Example 2:

Input:
strs := []string{"dog", "racecar", "car"}

Output:
""

Explanation:
There is no common prefix among the input strings.

Example 3:

Input:
strs := []string{"", "b"}

Output:
""

Constraints:
- All strings consist of only lowercase English letters.
*/

package main

import (
	"fmt"
)

// longest_common_prefix: Implement your solution here
func longest_common_prefix(str []string) string {
	lct := ""

	prefix := str[0]

	for idx, ch := range prefix {
		fmt.Println(string(ch))
	}

	return lct
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	lcp := longest_common_prefix(strs)
	fmt.Printf("Longest common prefix is %s\n", lcp)
}
