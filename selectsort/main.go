package main

import "fmt"

func selsort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		minIdx := j
		i := j + 1
		for i < len(arr) {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
			i++
		}
		if minIdx != j {
			swpbuff := arr[j]
			arr[j] = arr[minIdx]
			arr[minIdx] = swpbuff
		}
	}
}

func main() {
	// arr := []int{5, 2, 4, 6, 1, 3}
	arr := []int{31, 41, 59, 26, 41, 58}
	fmt.Println(arr)
	selsort(arr)
	fmt.Println(arr)
}
