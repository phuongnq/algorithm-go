package main

import "fmt"

func main() {
	arr := []int{14, 33, 27, 10, 35, 19, 42, 44}
	fmt.Println(selectionSort(arr))
}

/**
  Step 1 − Set MIN to location 0
  Step 2 − Search the minimum element in the list
  Step 3 − Swap with value at location MIN
  Step 4 − Increment MIN to point to next element
  Step 5 − Repeat until list is sorted
*/
func selectionSort(arr []int) []int {
	var minVal, minIndex int
	for i := 0; i < len(arr); i++ {
		minVal = arr[i]
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minVal {
				minVal = arr[j]
				minIndex = j
			}
		}

		arr[minIndex] = arr[i]
		arr[i] = minVal
	}

	return arr
}
