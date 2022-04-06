/*
i: array with mix of numbers and arrays as values
o: print to console every value that is a number (no arrays)

r: don't miss any values, but don't print arrays. Use recursion.

algorithm:
write function that takes an array as argument
- iterate through the array
  - if the value is a number -> print to console
  - if the value is an array -> call the function and pass in the array
*/

const arr = [ 1,
  2,
  3,
  [4, 5, 6],
  7,
  [8,
  [9, 10, 11,
  [12, 13, 14]
  ]
  ],
  [15, 16, 17, 18, 19,
  [20, 21, 22,
  [23, 24, 25,
  [26, 27, 29]
  ], 30, 31
  ], 32
  ], 33
  ]

function printAllNums(arr) {
  arr.forEach(val => {
    if (Array.isArray(val)) {
      printAllNums(val)
    } else {
      console.log(val)
    }
  })
}

printAllNums(arr)