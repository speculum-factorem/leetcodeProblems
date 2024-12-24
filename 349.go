// package main

// import "fmt"

// func main() {

// 	var varic1 = []int{1, 2, 3, 4}
// 	var varic2 = []int{1, 2, 3}

// 	fmt.Println(intersection(varic1, varic2))

// }

// func intersection(nums1 []int, nums2 []int) []int {

// 	elMap := make(map[int]bool)

// 	for _, num := range nums1 {

// 		elMap[num] = true
// 	}

// 	result := []int{}

// 	for _, num := range nums2 {

// 		if _, value := elMap[num]; value {

// 			result = append(result, num)

// 			delete(elMap, num)
// 		}

// 	}

// 	return result

// }
