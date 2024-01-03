package main
import "fmt"
func main() {
	var num int
	fmt.Scan(&num)
	hour := num/3600
	num = num - hour * 3600
	minut := num/60
	fmt.Print("It is ", hour, " hours ", minut, " minutes.")
}