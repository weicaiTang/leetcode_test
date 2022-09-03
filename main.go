package main

import "fmt"

func main() {
	var nums = []int{1, 2, 45, 5, 6, 7, 0, 12}
	i := twoSum(nums, 8)
	fmt.Printf("i: %v\n", i)
}
