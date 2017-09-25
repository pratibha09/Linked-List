package main

import (
	"fmt"
)

func search([]arr int, n int, x int) int {
	for i := 0; i < n; i++ {
		if arr[i] == x {
			return i
		}
		return -1
	}
}
func main(){
	arr := make([]int, 5)
	a[0], a[1], a[2], a[3], a[4] = {1, 10, 20, 34, 45}
	var x int:= 30
	var n := len(arr)
	fmt.Printf("%d is present at index %d\n", x, search(arr, n, x))
}



