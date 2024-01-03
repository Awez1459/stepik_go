package main
import "fmt"
func main() {
    var big, small string
    fmt.Scan(&big, &small)
	var index int = -1
	for i := 0; i <= len(big)-len(small); i++ {
		found := true
		for j := 0; j < len(small); j++ {
			if big[i+j] != small[j] {
				found = false
				break
			}
		}
		if found {
			index = i
			break
		}
	}
	fmt.Print(index)
}