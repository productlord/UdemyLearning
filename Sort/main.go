package main

import (
	"fmt"
	"sort"
)
func max(numbers ...int) int {
	
	sort.Ints(numbers)

	return numbers[len(numbers)-1]

	}
	


func main() {
	greatest := max(4, 7, 9, 123, 543, 23, 435, 1153, 125, 8875)
	fmt.Println("greatest =", greatest)
}

/*
FYI
For your code to also work with only negative numbers such as

greatest := max(-200 -700)

include this as your range statement
for i, v := range numbers {
	if v > largest || i == 0 {
		largest = v
	}
}

What does that code do?

The first time through the range loop
the index, i, will be zero
so largest will be set to the first number

Originally largest is set to the zero value for an int, which is zero

Zero would be greater than any negative number

if you only have negative numbers
you need largest to be something less than zero

Thanks to Ricardo G for this code improvement!
*/
