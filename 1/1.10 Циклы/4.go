package main
import "fmt"
func main() {
	var number, max_number int
	fmt.Scan(&number)
	var count = 0
	max_number = number
	for number != 0 {
		if number > max_number{
			max_number = number
			count = 1
		}else if number == max_number {
			count += 1
		}
		fmt.Scan(&number)
	}
	fmt.Print(count)
}