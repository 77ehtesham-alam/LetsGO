package main

import "fmt"

// Bad variable naming and lack of error handling
func main() {
	x := 5
	y := 10
	z := 0

	// Bad looping structure
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			x += 2
		} else {
			y += 3
		}
	}

	// Bad error handling
	if z == 0 {
		fmt.Println("Error: Division by zero!")
	} else {
		result := x / z // Division by zero possible
		fmt.Println("Result:", result)
	}

	// No meaningful comments
	var sum int
	for i := 0; i < x; i++ {
		// Inefficiently calculating sum
		for j := 0; j < y; j++ {
			sum += i * j
		}
	}

	fmt.Println("Sum:", sum)
}
