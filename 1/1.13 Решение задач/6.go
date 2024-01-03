package main
import "fmt"
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := (a + b) / 2
	if (a+b)%2 == 0 {
		fmt.Print(result)
	} else {
		fmt.Print(float64(a+b) / 2.0)
	}
}