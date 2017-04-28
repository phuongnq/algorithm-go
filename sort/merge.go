package main

import (
	"fmt"
)

func main() {
	arr := []int{14, 33, 27, 10, 35, 19, 42, 44}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}

	halfIndex := n / 2
	left := arr[0:halfIndex]
	right := arr[halfIndex:n]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	lenA := len(a)
	lenB := len(b)
	i := 0
	j := 0
	k := 0
	c := make([]int, lenA+lenB)
	for i+j < lenA+lenB {
		if i < lenA && j < lenB {
			if a[i] < b[j] {
				c[k] = a[i]
				i = i + 1
				k = k + 1
			} else if a[i] > b[j] {
				c[k] = b[j]
				j = j + 1
				k = k + 1
			}
		} else if i < lenA {
			c[k] = a[i]
			i = i + 1
			k = k + 1
		} else if j < lenB {
			c[k] = b[j]
			j = j + 1
			k = k + 1
		}
	}

	return c
}
