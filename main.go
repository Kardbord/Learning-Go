package main

import (
	"fmt"
	"os"
	fib "scratch/fibonacci"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("This program takes any number of base-10 16-bit unsigned integer arguments.\n" +
			"For each of these arguments, the value of their place in the fibonacci sequence is calculated.")
		return
	}
	
	fmt.Println("Calculating your fibonacci numbers...")
	for argIdx, arg := range os.Args[1:] {
		intArg, err := strconv.ParseUint(arg, 10, 32)
		if err != nil || intArg == 0xFFFFFFFF {
			fmt.Printf("Passed in argument #%d = %s could not be converted to uint32.\n", argIdx, arg)
			continue
		}
		
		fmt.Printf("fibonacci[%d] = %d\n", intArg, fib.Fibonacci(uint32(intArg)))
	}
}
