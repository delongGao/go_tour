package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// I really don't like this implementation
// func fibonacci() func() int {
// 	var (
// 		low, high int;
// 		low_defined, high_defined bool;
// 	)
// 	return func() int {
// 		if !low_defined {
// 			low_defined = true
// 			low = 0
// 			return low
// 		} else if !high_defined {
// 			high_defined = true
// 			high = 1
// 			return high
// 		}

// 		if low < high {
// 			low += high
// 			return low
// 		} else {
// 			high += low
// 			return high
// 		}
// 	}
// }

// I know there will be better solutions

func fibonacci() func() int {
	// low, high, sum := 0, 1, 0

	// return func() int {
	// 	sum, low, high = low, high, low + high
	// 	return sum
	// }

	// NOTE: this is by far the version I like most
	// 	because the variable definition most most sense
	//  to me, and the use of defer is also reasonable
	current, next := 0, 1
	return func() int {
		defer func() {
			current, next = next, current+next
		}()
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
