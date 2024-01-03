package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	var sum = 0
	for i := 0; i < a; i++ {
		var numeric_a int
		fmt.Scan(&numeric_a)
		if numeric_a >= 10 && numeric_a < 100 && numeric_a%8==0 {
			sum+=numeric_a
		}
	}
	fmt.Print(sum)
}