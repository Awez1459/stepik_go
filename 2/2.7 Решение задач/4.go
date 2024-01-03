package main
import (
	"fmt"
	"strconv"
)
func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}
	for _, r := range text {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("invalid int value: %c", r))
		}
		fmt.Print(n * n)
	}
	fmt.Println()
}