/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	tmp := make(map[int]int)
	for k, v := range nums {
		index, ok := tmp[target-v]
		if ok && index != k {
			res[0] = index
			res[1] = k
		}
		tmp[v] = k
	}
	return res
}

func main() {
	nums := []int{3, 3}
	fmt.Println(twoSum(nums, 6))
	nums = []int{3, 2, 3}
	fmt.Println(twoSum(nums, 6))
}
