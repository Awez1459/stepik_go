package main
import "fmt"
func main(){
	var num int
	fmt.Scan(&num)
	for num <= 100 {
		if num>9 && num<=100{
			fmt.Print(num, "\n")
		}
		fmt.Scan(&num)
	}
}