// package main

// import "fmt"

// func main() {

// 	fmt.Println(findDuplicate([]int{1, 2, 3, 3}))
// }

// func findDuplicate(nums []int) int {

// 	numMap := make(map[int]bool)

// 	for _, el := range nums {
// 		if numMap[el] {
// 			return el
// 		} else {
// 			numMap[el] = true
// 		}
// 		fmt.Println(el)
// 	}
// 	return -1
// }
