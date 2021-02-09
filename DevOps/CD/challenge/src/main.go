package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"./calculator"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, request *http.Request) {
	aParam := request.URL.Query().Get("a")
	bParam := request.URL.Query().Get("b")

	aNumber, aErr := strconv.Atoi(aParam)
	bNumber, bErr := strconv.Atoi(bParam)

	if aErr != nil {
		fmt.Fprintf(w, "ERROR: %s", aErr.Error())
		return
	}

	if bErr != nil {
		fmt.Fprintf(w, "ERROR: %s", bErr.Error())
		return
	}

	var secretvalue string
	if len(os.Getenv("SECRET_VALUE")) > 0 {
		secretvalue = os.Getenv("SECRET_VALUE")
	} else {
		secretvalue = "Não configurado"
	}

	fmt.Fprintf(w, "%d + %d = %d\n", aNumber, aNumber, calculator.Sum(aNumber, bNumber))
	fmt.Fprintf(w, "CONFIG VALUE: %s\n", os.Getenv("MY_EXAMPLE_CONFIG_FIELD"))
	fmt.Fprintf(w, "SECRET VALUE: %s\n", secretvalue)

}
