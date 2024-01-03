package main
import "fmt"
func main() {
	var num, digit int
	fmt.Scan(&num, &digit)
	str1 := ""
	for num > 0 {
		if num % 10 != digit {
			str1 = fmt.Sprintf("%d%s", num%10, str1)
		}
		num /= 10
	}
	fmt.Print(str1)
}