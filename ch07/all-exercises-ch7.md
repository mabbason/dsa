Exercises
 The following exercises provide you with the opportunity to practice with algorithms in practical situations. The solutions to these exercises are found in the section,  Chapter 7  .
 # 1  Use Big O Notation to describe the time complexity of the following function. The function returns true if the array is a “100-Sum Array,” and false if it is not.
 A “100-Sum Array” meets the following criteria:
 •  Its first and last numbers add up to 100.
 •  Its second and second-to-last numbers add up to 100.
 •  Its third and third-to-last numbers add up to 100, and so on.
 Here is the function:
  def   one_hundred_sum? (array)
 left_index = 0
 right_index = array. length  - 1
 
  while  left_index < array. length  / 2
  if  array[left_index] + array[right_index] != 100
  return   false 
  end 
 
 left_index += 1
 right_index -= 1
  end 
 
  return   true 
  end 
  **response**
  This is essentially N/2 which reduces to O(N)
# 2 Use Big O Notation to describe the time complexity of the following function. It merges two sorted arrays together to create a new sorted array containing all the values from both arrays:
  def   merge (array_1, array_2)
 new_array = []
 array_1_pointer = 0
 array_2_pointer = 0
 //Run the loop until we've reached end of both arrays: 
  while  array_1_pointer < array_1. length  ||
 array_2_pointer < array_2. length 
 
   If we already reached the end of the first array, 
  // add item from second array: 
  if  !array_1[array_1_pointer]
 new_array << array_2[array_2_pointer]
 array_2_pointer += 1
  // If we already reached the end of the second array, 
  // add item from first array: 
  elsif  !array_2[array_2_pointer]
 new_array << array_1[array_1_pointer]
 array_1_pointer += 1
  // If the current number in first array is less than current 
  // number in second array, add from first array: 
  elsif  array_1[array_1_pointer] < array_2[array_2_pointer]
 new_array << array_1[array_1_pointer]
 array_1_pointer += 1
  // If the current number in second array is less than or equal 
  // to current number in first array, add from second array: 
  else 
 new_array << array_2[array_2_pointer]
 array_2_pointer += 1
  end 
  end 
 
  return  new_array
  end
**response**
Idk, I had a lot of trouble reading the algorithm. My sense is that it is O(N^2) bc it seems like looping through the entirety of two arrays. So N*M type situation, which in worst case scenario is pretty much N^2.

# 3 Use Big O Notation to describe the time complexity of the following function. This function solves a famous problem known as “finding a needle in the haystack.” Both the needle and haystack are strings. For example, if the needle is "def" and the haystack is "abcdefghi" , the needle is contained somewhere in the haystack, as "def" is a substring of "abcdefghi" . However, if the needle is "dd" , it cannot be found in the haystack of "abcdefghi" . This function returns true or false, depending on whether the needle can be found in the haystack:
  def   find_needle (needle, haystack)
 needle_index = 0
 haystack_index = 0
 
  while  haystack_index < haystack. length 
  if  needle[needle_index] == haystack[haystack_index]
 found_needle =  true 
 
  while  needle_index < needle. length 
  if  needle[needle_index] != haystack[haystack_index + needle_index]
 found_needle =  false 
  break 
  end 
 needle_index += 1
  end 
 
  return   true   if  found_needle
 needle_index = 0
  end 
 
 haystack_index += 1
  end 
 
  return   false 
  end 
**response**
This is N*M, where N is the haystack and M is the substring. So in worst case scenario N and M are the same length so O(N^2)
# 4 Use Big O Notation to describe the time complexity of the following function. This function finds the greatest product of three numbers from a given array:
  def   largest_product (array)
 largest_product_so_far = array[0] * array[1] * array[2]
 i = 0
 
  while  i < array. length 
 j = i + 1
  while  j < array. length 
 k = j + 1
  while  k < array. length 
  if  array[i] * array[j] * array[k] > largest_product_so_far
 largest_product_so_far = array[i] * array[j] * array[k]
  end 
 k += 1
  end 
 j += 1
  end 
 i += 1
  end
return  largest_product_so_far
  end
**response**
This has three solid loops over all the elements, so O(N^3)
# 5 I once saw a joke aimed at HR people: “Want to immediately eliminate the unluckiest people from your hiring process? Just take half of the resumes on your desk and throw them in the trash.”  If we were to write software that kept reducing a pile of resumes until we had one left, it might take the approach of alternating between throwing out the top half and the bottom half. That is, it will first eliminate the top half of the pile, and then proceed to eliminate the bottom half of what remains. It keeps alternating between eliminating the top and bottom until one lucky resume remains, and that’s who we’ll hire! Describe the efficiency of this function in terms of Big O:
def   pick_resume (resumes)
 eliminate =  "top" 
 
  while  resumes. length  > 1
  if  eliminate ==  "top" 
 resumes = resumes[resumes. length  / 2, resumes. length  - 1]
 eliminate =  "bottom" 
  elsif  eliminate ==  "bottom" 
 resumes = resumes[0, resumes. length  / 2]
 eliminate =  "top" 
  end 
  end 
 
  return  resumes[0]
  end 
This is a O(Log N) algorithm, for each step, N is reduced by half.