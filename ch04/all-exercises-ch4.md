Exercises
 The following exercises provide you with the opportunity to practice with speeding up your code. The solutions to these exercises are found in the section,  Chapter 4  .

 # 1  Replace the question marks in the following table to describe how many steps occur for a given number of data elements across various types of Big O:
 N Elements    O(N)    O(log N)    O(N^2)
 100           100     6           10000
 2000          2000    10          4000000 

 # 2  If we have an O(N^2) algorithm that processes an array and find that it takes 256 steps, what is the size of the array?
 **response**
 array length of 16

 # 3  Use Big O Notation to describe the time complexity of the following function. It finds the greatest product of any pair of two numbers within a given array:
  def   greatestProduct (array):
 greatestProductSoFar = array[0] * array[1]
 
  for  i, iVal  in  enumerate(array):
  for  j, jVal  in  enumerate(array):
  if  i != j  and  iVal * jVal > greatestProductSoFar:
 greatestProductSoFar = iVal * jVal
 
  return  greatestProductSoFar
 **response**
 O(N^2) bc for each element in the array, it iterates over the whole array again.

 # 4  The following function finds the greatest single number within an array, but has an efficiency of O(N^2). Rewrite the function so that it becomes a speedy O(N):
  def   greatestNumber (array):
  for  i  in  array:

//Assume for now that i is the greatest: 
 isIValTheGreatest = True
 
  for  j  in  array:
// If we find another value that is greater than i,
// i is not the greatest: 
  if  j > i:
 isIValTheGreatest = False
 
//If, by the time we checked all the other numbers, i 
// is still the greatest, it means that i is the greatest number: 
  if  isIValTheGreatest:
  return  i
 **response**
 def   greatestNumber (array):

  greatestNum = array[0]

  for  i  in  array:
    if array[i] > greatestNum
    greatestNum = array[i]

  return greatestNum

/*
This only iterates through the array once, and so has a time complexity of O(N).
*/