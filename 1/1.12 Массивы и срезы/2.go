package main
import "fmt"
func main() {
	var length int
	fmt.Scan(&length)
	arr := make([]int, length)
	if length>=4{
		for i := 0; i < length; i++ {
			fmt.Scan(&arr[i])
		}
	}
	fmt.Print(arr[3])
}