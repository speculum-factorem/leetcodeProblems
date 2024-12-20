// package plusOnePackage

// func PlusOne(digits []int) []int {
// 	n := len(digits)

// 	// Начинаем с последнего элемента и увеличиваем его на 1
// 	for i := n - 1; i >= 0; i-- {
// 		digits[i]++
// 		if digits[i] < 10 {
// 			return digits // Если нет переполнения, возвращаем результат
// 		}
// 		digits[i] = 0 // Если есть переполнение, устанавливаем текущую цифру в 0
// 	}

// 	// Если все цифры были 9, создаем новый массив
// 	newArr := make([]int, n+1)
// 	newArr[0] = 1 // Устанавливаем первую цифру в 1
// 	return newArr
// }
