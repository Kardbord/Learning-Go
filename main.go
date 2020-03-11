package main

import (
	"errors"
	"fmt"
	"os"
	fib "scratch/fibonacci"
	"strconv"
)

// ----------------------------------- GLOBALS ------------------------------------ //

type program struct {
	name        string
	helpMsg     string
	runFunction func()error
}

var knownPrograms = map[string]program {
	"fibonacci": {
		"fibonacci",
		"The 'fibonacci' program takes any number of base-10 16-bit unsigned integer arguments.\n" +
			"For each of these arguments, the value of their place in the fibonacci sequence is calculated.",
		runFib,
	},
}

// ------------------------------------- MAIN ------------------------------------- //

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided, printing help.")
		fmt.Println("This program takes the following as command line arguments:\n" +
			              "    $1  : a string specifying which program to run. Known programs are", getProgNames(), "\n" +
			              "    $2  : An argument to pass to the program specified in $1\n" +
			              "    ... : An argument to pass to the program specified in $1")
		return
	}
	
	if progToRun, ok := knownPrograms[os.Args[1]]; ok {
		err := progToRun.runFunction()
		if err != nil {
			fmt.Println(err)
			fmt.Println(progToRun.helpMsg)
		}
	} else {
		fmt.Printf("'%s' does not match any known functionality. Supported functions are %v", os.Args[1], getProgNames())
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

// ------------------------------------ HELPERS ----------------------------------- //

func getProgNames() []string {
	progNames := make([]string, len(knownPrograms), len(knownPrograms))
	for key := range knownPrograms {
		progNames = append(progNames, "'" + key + "' ")
	}
	return progNames
}
