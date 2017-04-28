package main

import "fmt"

func main() {
	arr := []int{14, 33, 27, 10, 35, 19, 42, 44}
	fmt.Println(bubleSort(arr))
}

func bubleSort(arr []int) []int {
	var tmp int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}
