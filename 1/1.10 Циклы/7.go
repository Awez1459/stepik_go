package main
import "fmt"
func main() {
	var x, p, y int
	years := 0
	fmt.Scan(&x, &p, &y)
	for x < y{
		x = x + x*p/100
		years++
	}
	fmt.Print(years)
}