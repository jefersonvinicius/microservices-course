package main


import (
	"fmt" 
	"net/http"
	"./calc"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := calc.LoopingSqrt()
		fmt.Fprintf(w, "Code.education Rocks, %f\n", x)
	})

	http.ListenAndServe(":8080", nil)
}