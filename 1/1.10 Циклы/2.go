package main
import "fmt"
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var c int = 0
	for i := a; i <= b; i++ {
		c += i
	}
	fmt.Print(c)
}