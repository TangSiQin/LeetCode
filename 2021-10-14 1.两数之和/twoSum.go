package main

import "fmt"

func main() {

	nums := []int{2,7,11,15}
	target := 17

	fmt.Print(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}

	for i, x := range nums {

		if p, ok := hashTable[target-x]; ok {
			return []int{i, p}
		}

		hashTable[x] = i
	}

	return nil
}