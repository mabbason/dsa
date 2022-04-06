package main

import "fmt"

func main() {
	str := "the quick brown box jumps over a lazy dog"
	fmt.Println(findMissingLetter(str))
}

//a=97 z=122

func findMissingLetter(str string) string {
	alphabet := []string{}
	for ascii := 97; ascii < 123; ascii++ {
		alphabet = append(alphabet, string(ascii))
	}

	strHash := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		char := ""
		if i == len(str)-1 {
			char = str[i:]
		} else {
			char = str[i : i+1]
		}
		strHash[char] = true
	}

	for _, value := range alphabet {
		if !strHash[value] {
			return value
		}
	}
	return ""
}
