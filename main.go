package main

import (
	"errors"
	"fmt"
	"os"
	fib "scratch/fibonacci"
	"scratch/fizzbuzz"
	"strconv"
)

// ----------------------------------- GLOBALS ------------------------------------ //

type program struct {
	name        string
	helpMsg     string
	runFunction func()error
}

var knownPrograms = map[string]program {
	"fizzbuzz": {
		"fizzbuzz",
		"The 'fizzbuzz' program optionally takes 3 base-10 integer arguments.\n" +
			"    $1 : The beginning of a range; must be less than $2\n" +
			"    $2 : The end of a range; must be greater than $1\n" +
			"    $3 : The size of the step to be taken while generating the range; must be greater than 0\n" +
			"If no arguments are provided, arguments are assumed to be as follows:\n" +
			"    $1 = 1\n" +
			"    $2 = 100\n" +
			"    $3 = 1\n" +
			"For each of these arguments, one of four things will occur.\n" +
			"    1. 'Fizz' will be printed if the argument is evenly divisible by 3\n" +
			"    2. 'Buzz' will be printed if the argument is evenly divisible by 5\n" +
			"    3. 'FizzBuzz' will be printed if the argument is evenly divisible by both 3 and 5\n" +
			"    4. The argument will be printed if it is not evenly divisible by 3 or 5",
		runFizzBuzz,
	},
	"fibonacci": {
		"fibonacci",
		"The 'fibonacci' program takes any number of base-10 16-bit unsigned integer arguments.\n" +
			"For each of these arguments, the value of their place in the fibonacci sequence is calculated.",
		runFib,
	},
}

// ------------------------------------- MAIN ------------------------------------- //

func main() {
	
	printHelp := func() {
		fmt.Println("This program takes the following as command line arguments:\n" +
			"    $1  : a string specifying which program to run. Known programs are", getProgNames(), "\n" +
			"    $2  : An argument to pass to the program specified in $1\n" +
			"    ... : An argument to pass to the program specified in $1")
	}
	
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided, printing help.")
		printHelp()
		return
	}
	helpFlags := []string{"-h", "--h", "-help", "--help", "help", "h"}
	if strListContains(os.Args[:2], helpFlags) {
		printHelp()
		return
	}
	
	if progToRun, ok := knownPrograms[os.Args[1]]; !ok {
		// Did not find the program in our map
		fmt.Printf("'%s' does not match any known functionality. Supported functions are %v", os.Args[1], getProgNames())
	} else {
		// Found the program in our map
		if strListContains(os.Args[2:], helpFlags) {
			fmt.Println(progToRun.helpMsg)
			return
		}
		err := progToRun.runFunction()
		if err != nil {
			fmt.Println(err)
			fmt.Println(progToRun.helpMsg)
		}
	}
}

// ----------------------------------- RUNNERS ------------------------------------ //

func runFib() error {
	if len(os.Args) < 3 {
		return errors.New("\nthe fibonacci program requires at least 1 argument\n")
	}
	
	fmt.Println("Calculating your fibonacci numbers...")
	for argIdx, arg := range os.Args[2:] {
		intArg, err := strconv.ParseUint(arg, 10, 32)
		if err != nil || intArg == 0xFFFFFFFF {
			fmt.Printf("Passed in argument #%d = %s could not be converted to uint32.\n", argIdx, arg)
			continue
		}
		fmt.Printf("fibonacci[%d] = %d\n", intArg, fib.Fibonacci(uint32(intArg)))
	}
	return nil
}

func runFizzBuzz() error {
	fbArgs := os.Args[2:]
	if len(fbArgs) == 0 {
		list, err := NewSlice(1, 100, 1)
		if err != nil { return err }
		return fizzbuzz.Fizzbuzz(list)
	} else if len(fbArgs) != 3 {
		return fmt.Errorf("fizzbuzz requires exactly 0 or 3 arguments; %d arguments provided", len(fbArgs))
	} else {
		start, err := strconv.Atoi(fbArgs[0])
		if err != nil { return err }
		end, err := strconv.Atoi(fbArgs[1])
		if err != nil { return err }
		step, err := strconv.Atoi(fbArgs[2])
		if err != nil { return err }
		list, err := NewSlice(start, end, step)
		if err != nil { return err }
		return fizzbuzz.Fizzbuzz(list)
	}
}

// ------------------------------------ HELPERS ----------------------------------- //

func getProgNames() []string {
	progNames := make([]string, len(knownPrograms), len(knownPrograms))
	for key := range knownPrograms {
		progNames = append(progNames, "'" + key + "' ")
	}
	return progNames
}

func NewSlice(start, end, step int) ([]int, error) {
	if step <= 0 {
		return nil, fmt.Errorf("step size cannot be zero or negative; provided step was %d", step)
	}
	if start > end {
		return nil, errors.New("'start' cannot be greater than 'end'")
	}
	slice := make([]int, 0, ((end - start) / step) + 1)
	
	for start <= end {
		slice = append(slice, start)
		start += step
	}
	
	return slice, nil
}

func strListContains(strlist []string, keys []string) bool {
	for _, val := range strlist {
		for _, key := range keys {
			if val == key { return true }
		}
	}
	return false
}
