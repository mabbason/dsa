Exercises
 The following exercises provide you with the opportunity to practice with Big O Notation. The solutions to these exercises are found in the section,  Chapter 3  .

 # 1  Use Big O Notation to describe the time complexity of the following function that determines whether a given year is a leap year:

  function  isLeapYear(year) {
  return  (year % 100 === 0) ? (year % 400 === 0) : (year % 4 === 0);
 }

 **response**
O(1) bc the number of steps stays constant regardless of the year that is input.

 # 2  Use Big O Notation to describe the time complexity of the following function that sums up all the numbers from a given array:

  function  arraySum(array) {
  let  sum = 0;
 
  for ( let  i = 0; i < array.length; i++) {
 sum += array[i];
 }
 
  return  sum;
 }
 **response**
 O(N) bc for each length increase of the input array, there is a commensurate increase in the number of steps taken.

 # 3  The following function is based on the age-old analogy used to describe the power of compounding interest:

 Imagine you have a chessboard and put a single grain of rice on one square. On the second square, you put 2 grains of rice, since that is double the amount of rice on the previous square. On the third square, you put 4 grains. On the fourth square, you put 8 grains, and on the fifth square, you put 16 grains, and so on.
 The following function calculates which square you’ll need to place a certain number of rice grains. For example, for 16 grains, the function will return 5, since you will place the 16 grains on the fifth square.

Use Big O Notation to describe the time complexity of this function, which is below:

  function  chessboardSpace(numberOfGrains) {
  let  chessboardSpaces = 1;
  let  placedGrains = 1;
 
  while  (placedGrains < numberOfGrains) {
 placedGrains *= 2;
 chessboardSpaces += 1;
 }
 
  return  chessboardSpaces;
 }
 **response**
 O(log N) bc everytime the input number doubles, the steps (in this case loops) increases by 1.

 # 4  The following function accepts an array of strings and returns a new array that only contains the strings that start with the character "a" . Use Big O Notation to describe the time complexity of the function:

  function  selectAStrings(array) {
  let  newArray = [];
 
  for ( let  i = 0; i < array.length; i++) {
  if  (array[i].startsWith( "a" )) {
 newArray.push(array[i]);
 }
 }
 
  return  newArray;
 }
**response**
O(N) bc the algorithm is essentially just iterating over the input array. The number of steps increases proportionally with the length of the input array.

 # 5  The following function calculates the median from an ordered array. Describe its time complexity in terms of Big O Notation:

  function  median(array) {
  const  middle = Math.floor(array.length / 2);
 
  // If array has even amount of numbers: 
  if  (array.length % 2 === 0) {
  return  (array[middle - 1] + array[middle]) / 2;
 }  else  {  // If array has odd amount of numbers: 
  return  array[middle];
 }
 }
 **response**
 O(1) bc as the length of the input array increases the number of steps does not change.

