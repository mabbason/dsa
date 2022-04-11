package main

import "fmt"

func main() {
	fmt.Println(numOfShortestPaths(7, 3))
}

/*
 1  Imagine the function you’re writing has already been implemented by someone else.
 2  Identify the subproblem of the problem.
 3  See what happens when you call the function on the subproblem and go from there.
*/

/*
i: grid, (two ints)
o: int, (number of paths)

rules: use recursion

*/

func numOfShortestPaths(cols, rows int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return numOfShortestPaths(cols-1, rows) + numOfShortestPaths(cols, rows-1)
}
