package main
import "fmt"
func main() {
	var number int
	fmt.Scan(&number)
	lastDigit := number / 10 % 10
	fmt.Println(lastDigit)
}