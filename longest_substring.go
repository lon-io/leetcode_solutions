package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	longest := ""
	current := ""

	for _, v := range s {
		vStr := string(v)

		// Slower (20ms)
		// split := strings.Split(current, vStr)
		// if len(split) > 1 {
		//     fmt.Printf("\nFound %s, Resetting %s to: %s", vStr, current, split[1])
		//     current = split[1]
		// }

		// Faster (12 ms)
		oldIndex := strings.Index(current, vStr)
		if oldIndex > -1 {
			fmt.Printf("\nFound %s at %d, Resetting %s to: %s", vStr, oldIndex, current, current[oldIndex+1:])
			current = current[(oldIndex + 1):]
		}

		current += vStr

		if len(current) > len(longest) {
			longest = current
		}
	}

	fmt.Printf("\nLongest String: %s", longest)
	return len(longest)
}

// Slower (24 ms)
// Possible explanation here: https://stackoverflow.com/a/29690063/4931825
// String keys require hashing
func lengthOfLongestSubstringSlidingWindow(s string) int {
	soln := 0
	seenMap := map[string]int{}
	i := 0

	for j, v := range s {
		vStr := string(v)

		k, ok := seenMap[vStr]
		if ok && k >= i {
			i = k + 1
			fmt.Printf("\nFound %s at %d, Resetting i to: %d, with j == %d", vStr, k, i, j)
		}

		seenMap[vStr] = j
		current := j - i + 1
		if current > soln {
			soln = current
		}
	}

	return soln
}

func main() {
	test1 := "a"      // 1
	test2 := "bbbbbb" // 1
	test3 := "pwwkew" // 3
	test4 := "dvdf"   // 3
	fmt.Printf("\n%d", lengthOfLongestSubstring(test1))
	fmt.Printf("\n%d", lengthOfLongestSubstring(test2))
	fmt.Printf("\n%d", lengthOfLongestSubstring(test3))
	fmt.Printf("\n%d", lengthOfLongestSubstring(test4))

	fmt.Println("")

	fmt.Printf("\n%d", lengthOfLongestSubstringSlidingWindow(test1))
	fmt.Printf("\n%d", lengthOfLongestSubstringSlidingWindow(test2))
	fmt.Printf("\n%d", lengthOfLongestSubstringSlidingWindow(test3))
	fmt.Printf("\n%d", lengthOfLongestSubstringSlidingWindow(test4))
}
