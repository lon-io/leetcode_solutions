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

		// Faster
        oldIndex := strings.Index(current, vStr)
        if oldIndex > -1 {
			fmt.Printf("\nFound %s at %d, Resetting %s to: %s", vStr, oldIndex, current, current[oldIndex + 1:])
            current = current[(oldIndex + 1):]
        }

        current += string(vStr)

        if len(current) > len(longest) {
            longest = current
        }
    }

	fmt.Printf("\nLongest String: %s", longest)
    return len(longest)
}

func main() {
	test1 := "pwwkew"
	test2 := "dvdf"
	fmt.Printf("\n%d", lengthOfLongestSubstring(test1))
	fmt.Printf("\n%d", lengthOfLongestSubstring(test2))
}
