// package main

// import "fmt"

// func main() {

// 	randArr := []int{1, 2, 5}

// 	fmt.Println(searchInsert(randArr, 6))
// }

// func searchInsert(nums []int, target int) int {
// 	left := 0
// 	right := len(nums) - 1

// 	for left <= right {
// 		middle := left + (right-left)/2

// 		if nums[middle] == target {
// 			return middle // Найдено целевое значение
// 		} else if nums[middle] < target {
// 			left = middle + 1 // Ищем в правой половине
// 		} else {
// 			right = middle - 1 // Ищем в левой половине
// 		}
// 	}

// 	// Если не нашли, `left` указывает на место вставки
// 	return left
// }
