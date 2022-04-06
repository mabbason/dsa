Exercises
 The following exercises provide you with the opportunity to practice with optimizing for best- and worst-case scenarios. The solutions to these exercises are found in the section,  Chapter 6  .
 # 1  Use Big O Notation to describe the efficiency of an algorithm that takes 3N^2 + 2N + 1 steps.
**response**
O(N^2) bc Big O notation only factors in the largest exponent, and then ignores constants.
 # 2  Use Big O Notation to describe the efficiency of an algorithm that takes N + log N steps.
 **response**
 O(N) bc same as first response and `logN` is less steps than just N
 # 3  The following function checks whether an array of numbers contains a pair of two numbers that add up to 10.
  function  twoSum(array) {
  for  ( let  i = 0; i < array.length; i++) {
  for  ( let  j = 0; j < array.length; j++) {
  if  (i !== j && array[i] + array[j] === 10) {
  return   true ;
 }
 }
 }
  return   false ;
 }
 What are the best-, average-, and worst-case scenarios? Then, express the worst-case scenario in terms of Big O Notation.
 **response**
 Best case scenario is that index 0 and 1 have numbers that add up to 10. Worst case is that there aren't any numbers that add up to 10 and every single index is checked along with all the remaining indexes. Average is probably that somewhere around half of the array is checked and the algorithm returns early. The Big O notation for the worst case scenario is: O(N^2)
 # 4  The following function returns whether or not a capital “X” is present within a string.
  function  containsX(string) {
 
 foundX =  false ;
 
  for ( let  i = 0; i < string.length; i++) {
  if  (string[i] ===  "X" ) {
 foundX =  true ;
 }
 }
 
  return  foundX;
  }
 What is this function’s time complexity in terms of Big O Notation?
 Then, modify the code to improve the algorithm’s efficiency for best- and average-case scenarios.
 **response**
 O(N), basically a single pass through for every char. It would improve this algorithm to just move the return to within the `if block` and then as soon as a match is found (ideally asap) then the algorithm can return early instead of continuing for all of the needless iterations of checking the rest of the string.
