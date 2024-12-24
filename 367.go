package main

import "fmt"

func main() {

	fmt.Println(isPerfectSquare(1))
}

func isPerfectSquare(num int) bool {

	for i := 1; i <= (num+1)/2; i++ {
		if i*i == num {
			return true
		}
	}

	return false

}
