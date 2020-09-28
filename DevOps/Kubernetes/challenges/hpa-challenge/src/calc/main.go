package calc


import (
	"math"
)

func LoopingSqrt() float64 {
	x := 0.00001
	for i := 0; i < 9000000; i++ {
		x += math.Sqrt(x)  
	}
	return x;
}