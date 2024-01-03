package main
import "fmt"
func main() {
	var num int
	fmt.Scan(&num)
	count := 0
	for i := 0; i < num; i++ {
		var in_num int
		fmt.Scan(&in_num)
		if in_num == 0{
			count+=1
		}
	}
	fmt.Print(count)
}