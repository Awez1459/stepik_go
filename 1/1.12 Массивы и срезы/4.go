package main
import "fmt"
func main() {
	var length int
	fmt.Scan(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < length; i++ {
		if i%2==0{
			fmt.Print(arr[i], " ")
		}
	}
}