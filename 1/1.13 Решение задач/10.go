package main
import "fmt"
func main () {
	var a, b, maks int
	fmt.Scan(&a, &b)
	found := false
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			maks = i
			found = true
		}
	}
	if found {
		fmt.Print(maks)
	}else{
		fmt.Print("NO")
	}
}