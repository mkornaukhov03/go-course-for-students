package fizzbuzz_test

import (
	"github.com/stretchr/testify/assert"
	"lecture01_homework/fizzbuzz"
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var res string

	res = fizzbuzz.FizzBuzz(9)
	assert.Equal(t, res, "Fizz")

	res = fizzbuzz.FizzBuzz(25)
	assert.Equal(t, res, "Buzz")

	res = fizzbuzz.FizzBuzz(30)
	assert.Equal(t, res, "FizzBuzz")

	res = fizzbuzz.FizzBuzz(11)
	assert.Equal(t, res, "11")

	res = fizzbuzz.FizzBuzz(26)
	assert.Equal(t, res, "26")

	res = fizzbuzz.FizzBuzz(31)
	assert.Equal(t, res, "31")
}

func TestFizzBuzzBF(t *testing.T) {
	const (
		Begin   = 0
		IterCnt = 1000
		Step    = 1
	)

	var res string
	for i := Begin; i < IterCnt; i += Step {
		res = fizzbuzz.FizzBuzz(i)
		if i%15 == 0 {
			assert.Equal(t, res, "FizzBuzz")
		} else if i%3 == 0 {
			assert.Equal(t, res, "Fizz")
		} else if i%5 == 0 {
			assert.Equal(t, res, "Buzz")
		} else {
			assert.Equal(t, res, strconv.Itoa(i))
		}
	}
}
