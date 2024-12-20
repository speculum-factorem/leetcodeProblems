package main

import "fmt"

func main() {

	var arr = []int{1, 2, 3, 4, 5, 6}
	var avv int

	avv = 2

	fmt.Println(binarySearch(arr, avv))

}

func binarySearch(nums []int, target int) int {

	var left int
	var right int
	var middle int

	left = 0
	right = len(nums) - 1

	for left <= right {

		middle = left + (right-left)/2

		if nums[middle] == target {

			return middle

		} else if target > nums[middle] {

			left = middle + 1

		} else {

			right = middle - 1

		}
	}

	return -1
}
