package main
import "fmt"
func main() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		panic("invalid input values")
	}
	bs := []rune(text)
	maxR := bs[0]
	for i := 1; i < len(bs); i++ {
		r := bs[i]
		if r > maxR {
			maxR = r
		}
	}
	fmt.Println(string(maxR))
}