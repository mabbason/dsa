package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstNonDup("minimum"))
}

/*
i: single string
o: single char string
	- first NON-duplicated char

example:
i: "minimum"
o: "n" (m and i are both duplicated)

algorithm
write func takes single string as arg, returns string
- declare empty map hash table
- iterate through the string
	- for each char
		- if exists in hash table, increment occurences
		- else init in hash table w value of 1
- iterate through the string again
	- look up char in hash table
		- if value is 1, return char
		- else continue loop
- return emtpy string
*/

func firstNonDup(str string) string {
	charMap := make(map[string]int)

	for _, r := range str {
		char := string(r)
		if charMap[char] > 0 {
			charMap[char] += 1
		} else {
			charMap[char] = 1
		}
	}

	for _, r := range str {
		char := string(r)
		if charMap[char] == 1 {
			return char
		}
	}

	return ""
}
