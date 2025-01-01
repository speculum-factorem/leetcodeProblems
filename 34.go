// package main

// import "fmt"

// func main() {
// 	varic := []int{5, 7, 7, 8, 8, 10}
// 	fmt.Println(searchRange(varic, 8))
// }

// func searchRange(nums []int, target int) []int {

// 	left := 0
// 	right := len(nums) - 1
// 	middle := 0

// 	for left <= right {

// 		middle = left + (right-left)/2

// 		if middle > target && right != target {
// 			right = middle - 1
// 		} else if middle < target && left != target {
// 			left = middle + 1
// 		}

// 		if nums[right] == target && nums[left] == target {

// 			result := []int{left, right}
// 			return result

// 		}
// 	}

// 	result := []int{-1, -1}
// 	return result

// }
