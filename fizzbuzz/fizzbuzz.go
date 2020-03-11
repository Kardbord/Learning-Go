package fizzbuzz

import (
	"errors"
	"fmt"
)

func Fizzbuzz(nums []int) error {
	if nums == nil { return errors.New("passed in argument is nil") }
	for _, i := range nums {
		fizz := (i % 3) == 0
		buzz := (i % 5) == 0
		switch {
			case fizz && buzz: fmt.Println("FizzBuzz")
			case fizz:         fmt.Println("Fizz")
			case buzz:         fmt.Println("Buzz")
			default:           fmt.Println(i)
		}
	}
	return nil
}
