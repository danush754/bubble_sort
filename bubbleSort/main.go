package main

import "fmt"

func main() {

	var arr = []int{4, 3, 78, 2, 0, 2}

	bubble_sort(arr)
}

func bubble_sort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			var temp int
			if arr[j] > arr[i] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp

			}
		}
	}

	fmt.Println("arr", arr)
}
