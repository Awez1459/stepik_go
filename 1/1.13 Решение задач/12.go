package main
import "fmt"
func main () {
	var num int
	fmt.Scan(&num)
	for i := 1; i <= num; i*=2 {
		fmt.Print(i, " ")
	}
}