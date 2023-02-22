package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

const (
	Begin   = 0
	IterCnt = 100
	Step    = 1
)

func main() {
	for i := Begin; i < IterCnt; i += Step {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
