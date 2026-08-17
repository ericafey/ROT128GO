# ROT128 in Go

## Implementation of rot128
ROT128 rotates a byte by 128, which is the same as adding 1 bit 128 times with wraparound. 128 is the bit string `10000000` , meaning that adding this to any other 8-bit string flips the 8th bit. e.g.:
`01010101 + 10000000 = 11010101`
`11010101 + 10000000 = 101010101`
If the 8th bit of a byte is `1`, then adding 128 has it overflow into a 9th bit. We can use XOR to prevent this, whilst still flipping the 8th bit, e.g.:
 `11010101 XOR 10000000 = 01010101`

## Implementation Process
For writing the code I started by looking into Go a bit, such as looking at the specification and installing Go. I also wrote out some examples of rot128 to understand how it worked and to think about how best to implement it. Then I wrote a basic outline for the program and experimented a bit: reading input, rot128 implementation and printing the result. This did require a bit of reading for me as it was my first time working with Go. After this I wrote some (benchmarking) tests and looked into improving the performance. I tried not to spend too much time on creating too many tests or creating perfect code, as the assignment said I shouldn't. 

 ## Design Choices
 I tried to make the code minimal and readable, for example by using descriptive variable and function names, and thus also did not put many comments as it should hopefully be clear from the code. I also created separate functions to make the code readable and easy to test, namely a function for: rotating one byte at a time; a function that calls this function for every byte; and a function that reads input. To be precise and to prevent security issues, I made sure to make the type for variables clear by specifying []byte and byte.

## Tests
I wrote a few simple tests for the ROT128 logic, testing some happy paths and some edge cases. I focussed on making sure that it works for correct input, rather than malicious input, as a start.

## Performance Benchmarking
I looked into optimizing the rot128 implementation (not so much the input reading), but could not find too much to optimize. 
* In the `for` loop of the `RotateBytes` function, some memory allocation occurs because of the usage of slices; as the slice grows, new memory may need to be allocated, especially if the input is large. Making `rotatedBytes` an array would prevent this repeated allocation, however I did not find a good way to do this as you would need the input (constant) size to create the array.
* Parallelization seemed a bit out of scope for this assignment
* I thought of creating some sort of lookup table instead of performing the XOR operation every single time, but this operation is very small and quick, so I did not think it would be useful and it seemed a bit out of scope.