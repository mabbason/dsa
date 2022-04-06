// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

// 'abba' => true 
// acd => false 

func main() {
  // sumInts := []int{ 1, 2, 3 }

	// fmt.Println(sumRecursive(sumInts))
	// fmt.Println(prodFactorial(5))
	// isPalindrome("abba")
	fmt.Println(isPalindrome("abba")) // true
	fmt.Println(isPalindrome("abc")) // false
}


// func sumRecursive(slice []int) int{
// 	if len(slice) == 1 {
// 		return slice[0]
// 	}
// 	return slice[0] + sumRecursive(slice[1:])
// }

// func prodFactorial(num int) int {
// 	if num == 1 {
// 		return num
// 	}
// 	return num * prodFactorial(num - 1)
// }

/*
PALINDROME THOUGHTS
abccba
  cc  --> if middle two characters are the same it might be a palidrome, if not, it's not
 bccb --> if the subsequent next two inner characters are the same, same deal
abccba --> same as above

Rules:
	* always even number length string
	* runes will always be lowercase

ALGO
* determine if the middle two characters are a palindrome using magic, already-prepared `isPalindrome`; if it's not return false
* do the same with the next two outward characters
* return true

*/

func isPalindrome(input string) bool {
	if len(input) == 2 {
		return input[0] == input[1]
	}
	return isPalindrome(string(input[0]) + string(input[len(input) - 1])) && 
		isPalindrome(input[1:len(input)-1])
}
