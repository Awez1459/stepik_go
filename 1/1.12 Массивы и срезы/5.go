package main
import "fmt"
func main() {
	var length int
	fmt.Scan(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}
	count := 0
	for i := 0; i < length; i++ {
		if arr[i]>=0{
			count += 1
		}
	}
	fmt.Print(count)
}