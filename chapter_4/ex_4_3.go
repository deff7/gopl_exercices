package main

const size = 5

func reverse(arr *[size]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
