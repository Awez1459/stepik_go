package main
import "fmt"
func main() {
	var num int
	fmt.Scan(&num)
	count := 0
	for num > 0 {
		count += num%10
		num/=10
		if count >= 10 && num%10 == 0 {
			num = count
			count = 0
		}
	}
	fmt.Print(count)
}