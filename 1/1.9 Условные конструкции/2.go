package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	a1 := a / 1 % 10
	a2 := a / 10 % 10
	a3 := a / 100 % 10
	if a1 != a2 && a2 != a3 && a1 != a3 {
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}
}