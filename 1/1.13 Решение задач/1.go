package main
import "fmt"
func main() {
	var num int
	fmt.Scan(&num)
	count := 0
	for num > 0 {
		count += num%10
		num /= 10
	}
	fmt.Print(count)
}