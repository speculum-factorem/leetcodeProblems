// package main

// func main() {

// }

// func generateMatrix(n int) [][]int {

// 	matrix := make([][]int, n)

// 	for index := range matrix {
// 		matrix[index] = make([]int, n)
// 	}

// 	left, right, top, bottom := 0, len(matrix)-1, 0, len(matrix)-1
// 	currentNum := 1

// 	for left <= right && top <= bottom {

// 		for i := left; i <= right; i++ {
// 			matrix[top][i] = currentNum
// 			currentNum++
// 		}
// 		top++

// 		for i := top; i <= bottom; i++ {
// 			matrix[i][right] = currentNum
// 			currentNum++
// 		}
// 		right--

// 		if top <= bottom {
// 			for i := right; i >= left; i-- {
// 				matrix[bottom][i] = currentNum
// 				currentNum++
// 			}
// 		}
// 		bottom--

// 		if left <= right {
// 			for i := bottom; i >= top; i-- {
// 				matrix[i][left] = currentNum
// 				currentNum++
// 			}
// 			left++
// 		}
// 	}

// 	return matrix
// }
