package main
import (
	"fmt"
	"math"
)
func main() {
	var input float64
	fmt.Scan(&input)
	if input <= 0 {
		fmt.Printf("число %.2f не подходит\n", input)
	} else if input > 10000 {
		fmt.Printf("%e\n", input)
	} else {
		result := input * input
		result = math.Round(result*1e4) / 1e4 // Округляем до 4 знаков после запятой
		fmt.Printf("%.4f\n", result)
	}
}