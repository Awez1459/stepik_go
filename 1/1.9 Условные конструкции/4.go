package main
import "fmt"
func main() {
	var a int
	fmt.Scan(&a)
	a1 := a / 1 % 10
	a2 := a / 10 % 10
	a3 := a / 100 % 10
	a4 := a / 1000 % 10
	a5 := a / 10000 % 10
	a6 := a / 100000 % 10
	if a1+a2+a3==a4+a5+a6 {
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}
}