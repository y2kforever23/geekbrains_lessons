package fizzbuff

import "fmt"

func Fizzbuff() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

	}
}
