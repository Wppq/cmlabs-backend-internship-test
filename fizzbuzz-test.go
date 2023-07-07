package main

import "fmt"

func FizzBuzz(num int) string {
	fizz := num%3 == 0
	buzz := num%5 == 0

	if fizz && !buzz {
		return "Fizz"
	} else if !fizz && buzz {
		return "Buzz"
	} else if fizz && buzz {
		return "FizzBuzz"
	}

	return "can not proccess"
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i), "=>>>>", i)
	}
}
