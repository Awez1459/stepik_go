package main
import (
	"fmt"
)
func findFibonacciPosition(num int) int {
	a, b := 0, 1
	position := 0
	for b <= num {
		if b == num {
			return position
		}
		a, b = b, a+b
		position++
	}
	return -1
}
func main() {
	var fibNum int
	fmt.Scan(&fibNum)
	position := findFibonacciPosition(fibNum)
	if position != -1 {
		fmt.Println(position+1)
	} else {
		fmt.Println(-1)
	}
}