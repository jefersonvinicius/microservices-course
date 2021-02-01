package main

import (
	"fmt"
	"os"

	"./calculator"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	fmt.Printf("5 + 5 = %d\n", calculator.Sum(5, 5))
	fmt.Printf("CONFIG VALUE: %s\n", os.Getenv("MY_EXAMPLE_CONFIG_FIELD"))
}
