/**
*  insertion sort, solves the sorting problem
*  Input: A sequence of nnumbers⟨a1,a2,...,an⟩
* Output: A permutation (reordering) ⟨a1′ , a2′ , . . . , an′ ⟩ of the input sequence such
that a1′ ≤ a2′ ≤···≤ an′
*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{14, 33, 27, 10, 35, 19, 42, 44}
	fmt.Println(insertionSort(arr))

	arr = []int{3, 2, 7, 10}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	var position, val int
	for i := 0; i < len(arr); i++ {
		val = arr[i]
		position = i
		for position > 0 && arr[position-1] > val {
			arr[position] = arr[position-1]
			position = position - 1
		}
		arr[position] = val
	}

	return arr
}
