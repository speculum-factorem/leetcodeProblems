// package main

// import (
// 	"fmt"
// )

// func main() {
// 	board := [][]byte{
// 		{'X', '.', '.', 'X'},
// 		{'.', '.', '.', 'X'},
// 		{'.', '.', '.', 'X'},
// 	}

// 	count := countBattleships(board)
// 	fmt.Println(count) // Вывод количества кораблей
// }

// func countBattleships(board [][]byte) int {

// 	countOfBattleships := 0
// 	n := len(board)
// 	m := len(board[0])

// 	for i := 0; i < n; i++ {

// 		for j := 0; j < m; j++ {

// 			if board[i][j] == 'X' {
// 				countOfBattleships++
// 				dfs(board, i, j)
// 			}

// 		}
// 	}
// 	return countOfBattleships

// }

// func dfs(board [][]byte, i int, j int) {

// 	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '.' {
// 		return
// 	}

// 	board[i][j] = '.'

// 	dfs(board, i+1, j)
// 	dfs(board, i-1, j)
// 	dfs(board, i, j+1)
// 	dfs(board, i, j-1)

// }
