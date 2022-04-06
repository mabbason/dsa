package main

import "fmt"

//stack

type Stack struct {
	chars []string
}

//push
func (s *Stack) Push(str string) {
	s.chars = append(s.chars, str)
}

//pop
func (s *Stack) Pop() string {
	len := len(s.chars)

	popped := ""
	if len != 0 {
		popped = s.chars[len-1]
	}

	if len > 1 {
		s.chars = s.chars[0 : len-1]
	} else {
		s.chars = []string{}
	}
	return popped
}

/*
i: string, eg. "abcde"
o: same string, but reversed eg. "edcba"

rules:
use a stack in the implementation

algorithm
write func reverseString, takes single string as arg, returns string
- declare charsStack as empty stack
- iterate through the string
 - push each char onto stack
- declare empty string var as reversedStr
- while length of stack is not 0
	- pop off the stack and concat to reversedStr
return reversedStr
*/

func reverseString(str string) string {
	charsStack := Stack{}
	for _, c := range str {
		char := string(c)
		charsStack.Push(char)
	}
	reversedStr := ""
	for len(charsStack.chars) > 0 {
		popChar := charsStack.Pop()
		reversedStr += popChar
	}
	return reversedStr
}

func main() {
	fmt.Println(reverseString("abcde"))
}
