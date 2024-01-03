package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	if a >= 0 && a <= 9{
		fmt.Print(a)
	}else if a >= 10 && a <= 99{
		a1 := a / 10
		fmt.Print(a1)
	}else if a >= 100 && a <= 999{
		a1 := a / 100
		fmt.Print(a1)
	}else if a >= 1000 && a <= 9999{
		a1 := a / 1000
		fmt.Print(a1)
	}else if a >= 10000 && a <= 99999{
		a1 := a / 10000
		fmt.Print(a1)
	}
}