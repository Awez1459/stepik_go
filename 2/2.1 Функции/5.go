package main
import "fmt"
func sumInt(args ...int) (int, int) {
	sum := 0
	for _, x := range args {
		sum += x
	}
	return len(args), sum
}
func main() {
    var a, b int
    _, _ = fmt.Scan(&a, &b)
    fmt.Println(sumInt(a, b))
}