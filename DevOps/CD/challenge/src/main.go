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

	var secretvalue string
	if len(os.Getenv("SECRET_VALUE")) > 0 {
		secretvalue = os.Getenv("SECRET_VALUE")
	} else {
		secretvalue = "Não configurado"
	}

	fmt.Printf("CONFIG VALUE: %s\n", os.Getenv("MY_EXAMPLE_CONFIG_FIELD"))
	fmt.Printf("SECRET VALUE: %s\n", secretvalue)
}
