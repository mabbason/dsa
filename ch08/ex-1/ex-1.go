package main

import "fmt"

func main() {
	slc1 := []int{1, 2, 3, 4, 5, 7, 15, 13, 9, 9, 11}
	slc2 := []int{0, 2, 4, 6, 8, 7, 7, 7, 7}

	fmt.Println(arrIntersection(slc1, slc2))
}

func arrIntersection(slc1, slc2 []int) []int {
	if len(slc1) < len(slc2) {
		return getIntersectValues(slc1, slc2)
	} else {
		return getIntersectValues(slc2, slc1)
	}
}

func getIntersectValues(slcShort, slcLong []int) []int {
	hash := make(map[int]bool)
	for _, value := range slcLong {
		hash[value] = true
	}

	intersecting := []int{}
	for _, value := range slcShort {
		if hash[value] {
			if len(intersecting) == 0 {
				intersecting = append(intersecting, value)
			} else if intersecting[len(intersecting)-1] != value {
				intersecting = append(intersecting, value)
			}
		}
	}
	return intersecting
}
