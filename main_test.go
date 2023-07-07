package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizz(t *testing.T) {
	assert.Equal(t, "Fizz", FizzBuzz(3))
	assert.Equal(t, "Fizz", FizzBuzz(6))
	assert.Equal(t, "Fizz", FizzBuzz(9))
	assert.Equal(t, "Fizz", FizzBuzz(12))
	assert.Equal(t, "Fizz", FizzBuzz(36))
	assert.Equal(t, "Fizz", FizzBuzz(66))
	assert.Equal(t, "Fizz", FizzBuzz(99))
}

func TestBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", FizzBuzz(5))
	assert.Equal(t, "Buzz", FizzBuzz(10))
	assert.Equal(t, "Buzz", FizzBuzz(20))
	assert.Equal(t, "Buzz", FizzBuzz(25))
	assert.Equal(t, "Buzz", FizzBuzz(35))
	assert.Equal(t, "Buzz", FizzBuzz(40))
	assert.Equal(t, "Buzz", FizzBuzz(50))
	assert.Equal(t, "Buzz", FizzBuzz(55))
	assert.Equal(t, "Buzz", FizzBuzz(100))
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", FizzBuzz(15))
	assert.Equal(t, "FizzBuzz", FizzBuzz(30))
	assert.Equal(t, "FizzBuzz", FizzBuzz(45))
	assert.Equal(t, "FizzBuzz", FizzBuzz(60))
	assert.Equal(t, "FizzBuzz", FizzBuzz(75))
}

func TestValidation(t *testing.T) {
	assert.Equal(t, "can not proccess", FizzBuzz(2))
	assert.Equal(t, "can not proccess", FizzBuzz(4))
	assert.Equal(t, "can not proccess", FizzBuzz(8))
	assert.Equal(t, "can not proccess", FizzBuzz(38))
	assert.Equal(t, "can not proccess", FizzBuzz(64))
}
