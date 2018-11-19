package main

import "fmt"

func twoSum(nums []int, target int) []int {
    var numMap = map[int]int{}
    var result = []int{}
    for i, num := range nums {

        if j, ok := numMap[num]; ok {
            return []int{j, i};
        }

        var complement = target - num
        numMap[complement] = i
    }

    return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Printf("%#v\n", result)
}
