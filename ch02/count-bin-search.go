package main

import "fmt"

func main() {
	count := 0
	startNum := 100000
	for startNum > 2 {
		count++
		startNum /= 2
	}
	fmt.Println(count)
}
