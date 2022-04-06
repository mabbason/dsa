Exercises
 The following exercises provide you with the opportunity to practice analyzing algorithms. The solutions to these exercises are found in the section,  Chapter 5  .
 # 1  Use Big O Notation to describe the time complexity of an algorithm that takes 4N + 16 steps.
**response**
This would fall into the category of O(N) bc all the constants are dropped. Big O is only concerned with exponents
 # 2  Use Big O Notation to describe the time complexity of an algorithm that takes 2N^2 .
**response**
This would fall into the category of O(N^2) bc all the constants are dropped. Big O is only concerned with exponents
 # 3  Use Big O Notation to describe the time complexity of the following function, which returns the sum of all numbers of an array after the numbers have been doubled:
  def   double_then_sum (array)
 doubled_array = []
 
 array. each   do  |number|
 doubled_array << number *= 2
  end 
 
 sum = 0
 
 doubled_array. each   do  |number|
 sum += number
  end 
 
  return  sum
  end
**response**
This function has a time complexity of O(2N) which falls into the category of O(N).
 # 4  Use Big O Notation to describe the time complexity of the following function, which accepts an array of strings and prints each string in multiple cases:
  def   multiple_cases (array)
 array. each   do  |string|
 puts string. upcase 
 puts string. downcase 
 puts string. capitalize 
  end 
  end 
  This function has a time complexity of O(N) 
 # 5  The next function iterates over an array of numbers, and for each number whose index is even, it prints the sum of that number plus every number in the array. What is this function’s efficiency in terms of Big O Notation?
  def   every_other (array)
 array. each_with_index   do  |number, index|
  if  index. even? 
 array. each   do  |other_number|
 puts number + other_number
  end 
  end 
  end 
  end 
 This function has a time complexity of O(N^2)