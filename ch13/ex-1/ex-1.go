package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 5, 6, 3}
	// fmt.Println(quicksort(test, 0, len(test)-1))
	fmt.Println(quickSort(test, 0, 5))
	// prints
	// [0, 1, 2, 5, 6, 7]
}

/*
Given an array of positive numbers, write a function that returns the
greatest product of any three numbers. The approach of using three
nested loops would clock in at O(N3
), which is very slow. Use sorting to
implement the function in a way that it computes at O(N log N) speed.
(There are even faster implementations, but we’re focusing on using
sorting as a technique to make code faster.)

i: slice of positive numbers
o: single number

r:
output is the result of multiplying three largest numbers in slice
Big O must be NlogN
must sort the slice

example:
i: [1, 9, 23, 61, 4, 8, 102, 56, 13]
o: 102 * 61 * 56 = 348432

algorithm
write partition func, takes

*/
/* My implementation */
// func partition(slc []int, low, pivIdx int) ([]int, int) {
// 	high := pivIdx - 1
// 	// fmt.Println("low", low, "high", high)
// 	for {
// 		for slc[low] < slc[pivIdx] {
// 			// fmt.Println("low", low)
// 			low += 1
// 		}
// 		for slc[high] > slc[pivIdx] {
// 			// fmt.Println("high", high)
// 			high -= 1
// 		}
// 		fmt.Println("1", slc)
// 		fmt.Println("low", low, "high", high)
// 		if low >= high {
// 			break
// 		} else {
// 			slc[low], slc[high] = slc[high], slc[low]
// 			fmt.Println("2", slc)
// 			low += 1
// 		}
// 	}
// 	slc[low], slc[pivIdx] = slc[pivIdx], slc[low]
// 	fmt.Println("3", slc)
// 	return slc, low
// }

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

/* My implementation */
// func quicksort(slc []int, low, high int) []int {
// 	if low <= high {
// 		pivot := len(slc) - 1
// 		slc, pivot = partition(slc, low, pivot)
// 		fmt.Println("qs1")
// 		slc = quicksort(slc, low, pivot-1)
// 		fmt.Println("qs2")
// 		slc = quicksort(slc, pivot+1, high)
// 	}
// 	return slc
// }

func quickSort(arr []int, low, high int) []int {

	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
