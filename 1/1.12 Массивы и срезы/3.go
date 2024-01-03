package main
import "fmt"
func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
    fmt.Print(max(array[0], array[1], array[2], array[3], array[4]))
}