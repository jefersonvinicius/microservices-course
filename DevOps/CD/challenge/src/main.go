package main

import (
	"fmt"

	"./calculator"
)

func main() {
	fmt.Printf("5 + 5 = %d\n", calculator.Sum(5, 5))
}
