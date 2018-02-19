package main

import (
	"fmt"
	"math")
	

func fib(x uint64) uint64 {
	if x == 0 {
		return 0
	}
	
	if x < 3 {
		return x - 1
	}

	return fib(x-1) + fib(x-2)

}

func main() {
	
	for j:= 1; j < 51; j++ {
		fmt.Println(fib(uint64(j))
	}
	
}
