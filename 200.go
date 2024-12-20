// package main

// import "fmt"

// func main() {
// 	grid := [][]byte{
// 		{'1', '1', '0', '0'},
// 		{'1', '0', '0', '1'},
// 		{'0', '0', '1', '1'},
// 	}

// 	result := numIslands(grid)
// 	fmt.Println("Количество островов:", result)
// }

// // dfs - рекурсивная функция для поиска в глубину.
// func dfs(grid [][]byte, i int, j int) {
// 	// Проверяем выход за границы или наличие воды ('0').
// 	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
// 		return
// 	}

// 	// Помечаем текущую ячейку как посещённую, заменяя '1' на '0'.
// 	grid[i][j] = '0'

// 	// Рекурсивно вызываем dfs для соседних ячеек (вверх, вниз, влево, вправо).
// 	dfs(grid, i-1, j) // вверх
// 	dfs(grid, i+1, j) // вниз
// 	dfs(grid, i, j-1) // влево
// 	dfs(grid, i, j+1) // вправо
// }

// // numIslands подсчитывает количество островов в сетке.
// func numIslands(grid [][]byte) int {
// 	if len(grid) == 0 {
// 		return 0
// 	}

// 	rows := len(grid)
// 	cols := len(grid[0])
// 	islandCount := 0

// 	// Проходим по всей сетке.
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if grid[i][j] == '1' { // Если нашли землю.
// 				islandCount++   // Новый остров найден.
// 				dfs(grid, i, j) // Запускаем DFS для пометки острова.
// 			}
// 		}
// 	}

// 	return islandCount
// }
