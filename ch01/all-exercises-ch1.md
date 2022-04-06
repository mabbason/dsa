1. For array w 100 elements, how many steps would each of the following operations take?

a. Reading - 1 step, you have the index and it goes right to it
b. Searching - 100 steps, you have to iterate every index
c. Insertion at beginning - 101 steps, move every element back one, then insert
d. Insertion at end - 1 step, just insert at end
e. Deletion at begin - 100 steps, delete, then copy every of 99 element back one slot
f. Deletion at end - 1 steps, just delete last element

2. For array SET w 100 elements, how many steps would each of the following operations take?

a. Reading - 1 step, you have the index and go right to it
b. Searching - 100 steps, you have to iterate every index
c. Insertion at beginning - 201 steps, search make sure element is NOT in set already
    then copy every element over one slot, then insert
d. Insertion at end - 101, search every slot confirm unique, then insert at end
e. Deletion at begin - 100 steps, delete, then copy every of 99 element back one slot
f. Deletion at end - 1 steps, just delete last element

3. How many steps in N will it take to determine how many of a given element are contained within an array. Answer in 'N'.

For array of length N, it will take N steps. We need to search every index once.