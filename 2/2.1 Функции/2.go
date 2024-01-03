package main
import "fmt"
func minimumFromFour() int {
	var min int
	_, _ = fmt.Scan(&min)
	for i := 1; i < 4; i++ {
		var x int
		_, _ = fmt.Scan(&x)
		if x < min {
			min = x
		}
	}
	return min
}
func main() {
	fmt.Println(minimumFromFour())
}