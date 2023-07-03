package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, []int{}, &result)
	return result
}

func backtrack(nums []int, current []int, result *[][]int) {
	if len(current) == len(nums) {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for _, num := range nums {
		if contains(current, num) {
			continue
		}
		current = append(current, num)
		backtrack(nums, current, result)
		current = current[:len(current)-1]
	}
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{1, 2, 3}
	permutations := permute(numbers)
	fmt.Println("Все перестановки чисел", numbers, ":")
	for _, permutation := range permutations {
		fmt.Println(permutation)
	}
}
