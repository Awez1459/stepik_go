package main
import "fmt"
func main() {
	var N int
	fmt.Scan(&N)
	binaryRepresentation := ""
	for ; N > 0; N /= 2 {
		binaryRepresentation = fmt.Sprintf("%d%s", N%2, binaryRepresentation)
	}
	fmt.Println(binaryRepresentation)
}