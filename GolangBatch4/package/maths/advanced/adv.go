package advanced

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("1 init advanced")
}

func Square(a int) int {
	return a * a
}

func Sin(angle float64) float64 {
	return math.Sin(angle)
}
